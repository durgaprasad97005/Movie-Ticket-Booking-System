package services

import (
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/dto"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/models"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

// Create User
func (service *AuthService) Register(form dto.UserRequest) (*dto.UserResponse, error) {
	// Check whether there exists another user
	if _, err := service.repository.Count(form.Email); err != nil {
		return nil, err
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user model
	user := models.UserModel{
		Name: form.Name, 
		Email: form.Email, 
		Phone: form.Phone, 
		Password: string(passwordHash), 
		Role: form.Role, 
		IsActive: form.IsActive, 
	}
	user.DocumentTime.Defaults()

	// Insert user
	if err := service.repository.Insert(&user); err != nil {
		return nil, err
	}

	response := dto.UserResponse{}
	response.Id = user.Id
	response.Name = user.Name
	response.Email = user.Email
	response.Phone = user.Phone
	response.Role = user.Role
	response.IsActive = user.IsActive
	response.CreatedAt = user.CreatedAt
	response.UpdatedAt = user.UpdatedAt

	return &response, nil
}