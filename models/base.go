package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DocumentId struct {
	Id bson.ObjectID `json:"id" bson:"_id"`
}

type DocumentTime struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (documentTime *DocumentTime) Defaults() {
	now := time.Now()

	documentTime.CreatedAt = now
	documentTime.UpdatedAt = now
}
