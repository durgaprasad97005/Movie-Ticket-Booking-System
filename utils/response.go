package utils

import (
	"github.com/gofiber/fiber/v3"
)

func ErrorResponse(message string, err error) fiber.Map {
	return fiber.Map{
		"success": false,
		"message": message, 
		"error": err.Error(), 
	}
}

func SuccessResponse(message string, data any) fiber.Map {
	return fiber.Map{
		"success": true,
		"message": message, 
		"data": data, 
	}
}