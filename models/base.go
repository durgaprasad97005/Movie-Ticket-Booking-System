package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DocumentId struct {
	Id bson.ObjectID `json:"id" bson:"_id,omitempty"`
}

type CreatedAtTime struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type DocumentTime struct {
	CreatedAtTime `json:",inline" bson:",inline"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (documentTime *DocumentTime) Defaults() {
	now := time.Now()

	documentTime.CreatedAt = now
	documentTime.UpdatedAt = now
}
