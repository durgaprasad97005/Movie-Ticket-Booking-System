package dto

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/constants"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/models"
)

type UserRequest struct {
	Name     string                 `json:"name" validate:"required,min=3,max=25"`
	Email    string                 `json:"email" validate:"required,email"`
	Phone    string                 `json:"phone" validate:"required,e164"`
	Password string                 `json:"password" validate:"required,min=8,max=15"`
	Role     constants.UserRoleEnum `json:"role" validate:"required"`
	IsActive bool                   `json:"isActive"`
}

type UserResponse struct {
	models.DocumentId   `json:",inline" bson:",inline"`
	Name                string                 `json:"name"`
	Email               string                 `json:"email"`
	Phone               string                 `json:"phone"`
	Role                constants.UserRoleEnum `json:"role"`
	IsActive            bool                   `json:"isActive"`
	models.DocumentTime `json:",inline" bson:",inline"`
}
