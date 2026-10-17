package middleware

import "github.com/gofiber/fiber/v3"

func AuthMiddleware(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}