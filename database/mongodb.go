package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	UsersCollection = "users"
	RolesCollection = "roles"
	MoviesCollection = "movies"
	TheatresCollection = "theatres"
	ScreensCollection = "screens"
	SeatsCollection = "seats"
	ShowsCollection = "shows"
	BookingsCollection = "bookings"
	BookingSeatsCollection = "booking_seats"
	PaymentsCollection = "payments"
	ReviewsCollection = "reviews"
)

var DB *mongo.Database

func ConnectMongo(mongoUri string, dbName string) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Failed to ping to MongoDB: ", err)
	}

	DB = client.Database(dbName)
}

func Collection(coll string) *mongo.Collection {
	return DB.Collection(coll)
}