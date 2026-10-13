package routes

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/config"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/handlers"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/repositories"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/services"
	"github.com/gofiber/fiber/v3"
)

func NewAuthRoutes(app fiber.Router, cfg *config.Config) {
	// Initialize handler, service, and repository files
	repository := repositories.NewAuthRepository()
	service := services.NewAuthService(repository, cfg)
	handler := handlers.NewAuthHandler(service)

	// Initialize routes
	newApp := app.Group("/auth")
	newApp.Post("/register", handler.Register)
	newApp.Post("/login", handler.Login)
	newApp.Post("/refresh", handler.Refresh)
	newApp.Post("/logout", handler.Logout)
}