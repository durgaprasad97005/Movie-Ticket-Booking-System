package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type JWTRegreshTokenModel struct {
	DocumentId    `json:",inline" bson:",inline"`
	UserId        bson.ObjectID `bson:"userId"`
	TokenHash     string        `bson:"tokenHash"`
	ExpiresAt     time.Time     `bson:"expiresAt"`
	CreatedAtTime `bson:",inline"`
}
