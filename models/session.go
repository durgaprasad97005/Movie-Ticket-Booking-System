package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserSessionModel struct {
	DocumentId       `json:",inline" bson:",inline"`
	UserId           bson.ObjectID `bson:"userId"`
	RefreshTokenHash string        `bson:"refreshTokenHash"`
	ExpiresAt        time.Time     `bson:"expiresAt"`
	RevokedAt         *time.Time    `bson:"revokedAt"`
	CreatedAtTime    `bson:",inline"`
}
