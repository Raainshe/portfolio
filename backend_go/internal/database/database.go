package database

import (
	"backend_go/internal/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetCollection(name string) *mongo.Collection
	CreateUser(user models.User) (string, error)
	GetUser(userID string) (*models.User, error)
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
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))

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
		"tracking",
	}

	for _, coll := range collections {
		log.Printf("Collection: %s exists\n", coll)
	}
}

func (s *service) GetCollection(collName string) *mongo.Collection {
	return s.database.Collection(collName)
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
