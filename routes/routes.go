package routes

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App) {
	appGrp := app.Group("/api/v1")

	NewAuthRoutes(appGrp)
}