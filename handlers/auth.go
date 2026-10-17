package handlers

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/dto"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/services"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/utils"
	"github.com/go-playground/validator/v10"
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
func (handler *AuthHandler) Register(ctx fiber.Ctx) error {
	// parse the request
	var form dto.UserRequest
	if err := ctx.Bind().Body(&form); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid data", err),
		)
	}

	// validate the request payload
	validate := validator.New()
	if err := validate.Struct(&form); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid data", err),
		)
	}

	response, err := handler.service.Register(form)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("User creation failed", err), 
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		utils.SuccessResponse("User creation successful", response), 
	)
}

// Login user
func (handler *AuthHandler) Login(ctx fiber.Ctx) error {
	// parse the request
	var form dto.UserLoginRequest
	if err := ctx.Bind().Body(&form); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid data", err),
		)
	}

	// validate the request payload
	validate := validator.New()
	if err := validate.Struct(&form); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid data", err),
		)
	}

	response, err := handler.service.Login(form)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ErrorResponse("User login failed", err), 
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		utils.SuccessResponse("User login successful", response), 
	)
}