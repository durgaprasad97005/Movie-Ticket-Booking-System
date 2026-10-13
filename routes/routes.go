package routes

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/config"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	appGrp := app.Group("/api/v1")

	NewAuthRoutes(appGrp, cfg)
}