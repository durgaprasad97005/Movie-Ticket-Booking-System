package handlers

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/services"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(serv *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: serv,
	}
}

// Create user
func (handler *AuthHandler) Create(ctx fiber.Ctx) error {
	return ctx.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
	})
}