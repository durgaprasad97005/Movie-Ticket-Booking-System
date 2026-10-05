package repositories

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthRepository struct {
	collection *mongo.Collection // users collection
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		collection: database.Collection(database.UsersCollection),
	}
}

