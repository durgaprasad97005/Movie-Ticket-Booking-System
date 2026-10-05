package main

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/config"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func SetupApp(cfg *config.Config) *fiber.App {
	app := fiber.New()

	// Global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Health check api
	app.Get("/health", func(ctx fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"status": "ok"})
	})

	// Initialize routes
	routes.SetupRoutes(app)

	return app
}