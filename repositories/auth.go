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
	userSessionsCollection *mongo.Collection
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		collection: database.Collection(database.UsersCollection),
		userSessionsCollection: database.Collection(database.UserSessionsCollection),
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
	if user == nil {
		return errors.ErrInvalidCredentials
	}

	result, err := repository.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	user.Id = result.InsertedID.(bson.ObjectID)
	return nil
}

// Insert a single Token
func (repository *AuthRepository) InsertUserSession(token *models.UserSessionModel) error {
	if token == nil {
		return errors.ErrInvalidCredentials
	}

	result, err := repository.userSessionsCollection.InsertOne(context.TODO(), token)
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

// Get user session by refreshToken
func (repository *AuthRepository) GetUserSessionByRefreshToken(refreshToken string) (*models.UserSessionModel, error) {
	if refreshToken == "" {
		return nil, errors.ErrInvalidCredentials
	}

	var userSession models.UserSessionModel
	if err := repository.userSessionsCollection.FindOne(context.TODO(), bson.M{"refreshToken": refreshToken}).Decode(&userSession); err != nil {
		return nil, err
	}

	return &userSession, nil
}

// Update user session
func (repository *AuthRepository) UpdateUserSession(refreshToken string, fields bson.M) error {
	if refreshToken == "" {
		return errors.ErrInvalidCredentials
	}

	result, err := repository.userSessionsCollection.UpdateOne(context.TODO(), bson.M{"refreshToken": refreshToken}, bson.M{"$set": fields})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		return errors.ErrUserSessionNotFound
	}

	return nil
}