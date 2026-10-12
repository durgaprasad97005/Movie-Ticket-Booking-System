package models

import "github.com/durgaprasad97005/Movie-Ticket-Booking-System/constants"

type UserModel struct {
	DocumentId   `json:",inline" bson:",inline"`
	Name         string                 `bson:"name"`
	Email        string                 `bson:"email"`
	Phone        string                 `bson:"phone"`
	Password     string                 `bson:"password"`
	Role         constants.UserRoleEnum `bson:"role"`
	IsActive     bool                   `bson:"isActive"`
	DocumentTime `json:",inline" bson:",inline"`
}
