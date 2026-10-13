package services

import (
	"time"

	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/config"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/dto"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/models"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *repositories.AuthRepository
	config     *config.Config
}

func NewAuthService(repo *repositories.AuthRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		repository: repo,
		config:     cfg,
	}
}

// Create User
func (service *AuthService) Register(form dto.UserRequest) (*dto.JWTResponse, error) {
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
		Name:     form.Name,
		Email:    form.Email,
		Phone:    form.Phone,
		Password: string(passwordHash),
		Role:     form.Role,
		IsActive: form.IsActive,
	}
	user.DocumentTime.Defaults()

	// Insert user
	if err := service.repository.InsertUser(&user); err != nil {
		return nil, err
	}

	now := time.Now()

	// Refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id.Hex(),
		"iat": now.Unix(),
		"exp": now.Add(7 * 24 * time.Hour).Unix(),
	})

	jwtRefresh, err := refreshToken.SignedString(service.config.JWTSecret)
	if err != nil {
		return nil, err // user created but unable to create refresh token
	}

	// Access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id.Hex(),
		"iat": now.Unix(),
		"exp": now.Add(15 * time.Minute).Unix(),
	})

	jwtAccess, err := accessToken.SignedString(service.config.JWTSecret)
	if err != nil {
		return nil, err // user created but unable to create access token
	}

	// Create refreshTokens model
	token := models.JWTRegreshTokenModel{}
	token.UserId = user.Id
	token.TokenHash = jwtRefresh
	token.ExpiresAt = now.Add(7 * 24 * time.Hour)
	token.CreatedAt = now

	// Insert refresh token
	if err := service.repository.InsertToken(&token); err != nil {
		return nil, err
	}

	tokenResponse := dto.JWTResponse{}
	tokenResponse.RefreshToken = jwtRefresh
	tokenResponse.AccessToken = jwtAccess

	return &tokenResponse, nil
}
