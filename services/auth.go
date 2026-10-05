package services

import "github.com/durgaprasad97005/Movie-Ticket-Booking-System/repositories"

type AuthService struct {
	repository *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		repository: repo,
	}
}