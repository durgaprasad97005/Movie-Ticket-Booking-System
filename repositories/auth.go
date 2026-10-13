package repositories

import (
	"context"

	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/database"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/errors"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthRepository struct {
	collection *mongo.Collection // users collection
	refreshTokensCollection *mongo.Collection
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		collection: database.Collection(database.UsersCollection),
		refreshTokensCollection: database.Collection(database.RefreshTokensCollection),
	}
}

// Count documents
func (repository *AuthRepository) Count(mail string) (int64, error) {
	if mail == "" {
		return -1, errors.ErrInvalidCredentials
	}

	count, err := repository.collection.CountDocuments(context.TODO(), bson.M{"email": mail})

	return count, err
}

// Insert a single user
func (repository *AuthRepository) InsertUser(user *models.UserModel) error {
	result, err := repository.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	user.Id = result.InsertedID.(bson.ObjectID)
	return nil
}

// Insert a single Token
func (repository *AuthRepository) InsertToken(token *models.JWTRegreshTokenModel) error {
	result, err := repository.refreshTokensCollection.InsertOne(context.TODO(), token)
	if err != nil {
		return err
	}

	token.Id = result.InsertedID.(bson.ObjectID)
	return nil
}

// Get user by email
func (repository *AuthRepository) GetUserByEmail(mail string) (*models.UserModel, error) {
	if mail == "" {
		return nil, errors.ErrInvalidCredentials
	}

	var user models.UserModel
	if err := repository.collection.FindOne(context.TODO(), bson.M{"email": mail}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}