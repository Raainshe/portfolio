package database

import (
	"backend_go/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Health() map[string]string
	GetCollection(name string) *mongo.Collection
	CreateUser(user models.User) (string, error)
	GetUser(userID string) (*models.User, error)
	TrackVisitor(ipAddress, userAgent string) (*models.Visitor, error)
}

type service struct {
	db       *mongo.Client
	database *mongo.Database
}

var (
	host     = os.Getenv("BLUEPRINT_DB_HOST")
	port     = os.Getenv("BLUEPRINT_DB_PORT")
	database = os.Getenv("BLUEPRINT_DB_DATABASE")
)

func New() Service {
	username := os.Getenv("BLUEPRINT_DB_USERNAME")
	password := os.Getenv("BLUEPRINT_DB_ROOT_PASSWORD")

	//build a connections tring with authentication
	var uri string
	if username != "" && password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	log.Printf("Connecting to MongoDB with URI: %s", uri)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)

	}
	//create a database with a name
	databaseName := database
	if databaseName == "" {
		databaseName = "portfolio"
	}
	db := client.Database(databaseName)
	createCollections(db)
	return &service{
		db:       client,
		database: db,
	}
}

func createCollections(db *mongo.Database) {
	collections := []string{
		"users",
		"visitors",
	}

	for _, coll := range collections {
		log.Printf("Collection: %s exists\n", coll)
	}
}

func (s *service) GetCollection(collName string) *mongo.Collection {
	return s.database.Collection(collName)
}

func (s *service) TrackVisitor(ipAddress, userAgent string) (*models.Visitor, error) {
	collection := s.GetCollection("visitors")

	//check if they exist
	filter := bson.M{"ip_address": ipAddress}
	var existingVisitor models.Visitor
	err := collection.FindOne(context.Background(), filter).Decode(&existingVisitor)

	now := time.Now()

	if err == mongo.ErrNoDocuments {
		//this is a new visitor
		location, err := s.getLocationFromIP(ipAddress)
		if err != nil {
			log.Printf("could not get location ip %sL %v", ipAddress, err)
		}
		//continue even without location
		newVisitor := models.Visitor{
			IPAddress:  ipAddress,
			VisitCount: 1,
			VisitTimes: []time.Time{now},
			Location:   location,
			FirstVisit: now,
			LastVisit:  now,
			UserAgent:  userAgent,
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		result, err := collection.InsertOne(context.Background(), newVisitor)
		if err != nil {
			return nil, fmt.Errorf("failed to insert visitor: %w", err)
		}
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			newVisitor.ID = oid
		}

		return &newVisitor, nil
	} else if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	// now update old visitor

	existingVisitor.VisitCount++
	existingVisitor.VisitTimes = append(existingVisitor.VisitTimes, time.Now())
	existingVisitor.LastVisit = now
	existingVisitor.UpdatedAt = now

	//update location if it changed
	location, err := s.getLocationFromIP(ipAddress)
	if err == nil {
		existingVisitor.Location = location
	}

	update := bson.M{
		"$set": bson.M{
			"visit_count": existingVisitor.VisitCount,
			"visit_times": existingVisitor.VisitTimes,
			"last_visit":  existingVisitor.LastVisit,
			"updated_at":  existingVisitor.UpdatedAt,
			"location":    existingVisitor.Location,
		},
	}

	//update the user now
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update databse: %w", err)
	}

	return &existingVisitor, nil

}

func (s *service) getLocationFromIP(ipAddress string) (models.Location, error) {

	req := fmt.Sprintf("http://ip-api.com/json/%s", ipAddress)

	res, err := http.Get(req)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to call IP-API: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return models.Location{}, fmt.Errorf("IP-API returned status: %d", res.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return models.Location{}, fmt.Errorf("error decoding IP-API result: %w", err)
	}

	//check if api result was succesful
	if result["status"] != "success" {
		return models.Location{}, fmt.Errorf("IP-API returned error: %s", result["message"])
	}

	//extract location data
	location := models.Location{
		Country:     getString(result, "country"),
		CountryCode: getString(result, "countryCode"),
		Region:      getString(result, "region"),
		RegionName:  getString(result, "regionName"),
		City:        getString(result, "city"),
		Zip:         getString(result, "zip"),
		Lat:         getFloat(result, "lat"),
		Lon:         getFloat(result, "lon"),
		Timezone:    getString(result, "timezone"),
		ISP:         getString(result, "isp"),
	}

	return location, nil
}

// Helper functions for safe type conversion
func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getFloat(data map[string]interface{}, key string) float64 {
	if val, ok := data[key]; ok {
		if f, ok := val.(float64); ok {
			return f
		}
	}
	return 0
}
func (s *service) GetUser(userID string) (*models.User, error) {
	collection := s.GetCollection("users")

	//convert stringid to objectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id format: %w", err)
	}

	filter := bson.M{"_id": objID}

	//exclude specific fields
	projection := bson.M{
		"password": 0,
	}
	opts := options.FindOne().SetProjection(projection)

	var user models.User

	err = collection.FindOne(context.Background(), filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

func (s *service) CreateUser(user models.User) (string, error) {
	collection := s.GetCollection("users")

	//check if user already exists
	filter := bson.M{"email": user.Email}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return "", fmt.Errorf("databse error: %w", err)
	}
	if count > 0 {
		return "", fmt.Errorf("user already exists in database")
	}

	//hash password
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashpassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to get inserted user id")

}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
