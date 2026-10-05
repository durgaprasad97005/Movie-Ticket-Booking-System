package routes

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/handlers"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/repositories"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/services"
	"github.com/gofiber/fiber/v3"
)

func NewAuthRoutes(app fiber.Router) {
	// Initialize handler, service, and repository files
	repository := repositories.NewAuthRepository()
	service := services.NewAuthService(repository)
	handler := handlers.NewAuthHandler(service)

	// Initialize routes
	newApp := app.Group("/user")
	newApp.Post("/create", handler.Create)
}