package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Visitor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IPAddress  string             `bson:"ip_address" json:"ip_address"`
	VisitCount int                `bson:"visit_count" json:"visit_count"`
	VisitTimes []time.Time        `bson:"visit_times" json:"visit_times"`
	Location   Location           `bson:"location" json:"location"`
	FirstVisit time.Time          `bson:"first_visit" json:"first_visit"`
	LastVisit  time.Time          `bson:"last_visit" json:"last_visit"`
	UserAgent  string             `bson:"user_agent,omitempty" json:"user_agent,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type Location struct {
	Country     string  `bson:"country" json:"country"`
	CountryCode string  `bson:"country_code" json:"country_code"`
	Region      string  `bson:"region" json:"region"`
	RegionName  string  `bson:"region_name" json:"region_name"`
	City        string  `bson:"city" json:"city"`
	Zip         string  `bson:"zip" json:"zip"`
	Lat         float64 `bson:"lat" json:"lat"`
	Lon         float64 `bson:"lon" json:"lon"`
	Timezone    string  `bson:"timezone" json:"timezone"`
	ISP         string  `bson:"isp" json:"isp"`
}
