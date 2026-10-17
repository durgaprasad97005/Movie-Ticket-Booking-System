package services

import (
	"time"

	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/config"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/dto"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/errors"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/models"
	"github.com/durgaprasad97005/Movie-Ticket-Booking-System/repositories"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	count, err := service.repository.Count(form.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.ErrDuplicateUserFound
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user 
	user := models.UserModel{
		Name:         form.Name,
		Email:        form.Email,
		Phone:        form.Phone,
		PasswordHash: string(passwordHash),
		Role:         form.Role,
		IsActive:     form.IsActive,
	}
	user.DocumentTime.Defaults()

	if err := service.repository.InsertUser(&user); err != nil {
		return nil, err
	}

	// Create tokens
	refreshToken, accessToken, err := CreateTokens(user.Id, service.config.JWTSecret)
	if err != nil {
		return nil, err
	}
	if refreshToken == "" || accessToken == "" {
		return nil, errors.ErrInvalidTokens
	}

	now := time.Now()

	// Create session
	token := models.UserSessionModel{}
	token.UserId = user.Id
	token.RefreshTokenHash = refreshToken // need to store the hash instead of jwtRefresh token itself - future enhancement
	token.ExpiresAt = now.Add(7 * 24 * time.Hour)
	token.RevokedAt = nil
	token.CreatedAt = now

	if err := service.repository.InsertUserSession(&token); err != nil {
		return nil, err
	}

	tokenResponse := dto.JWTResponse{}
	tokenResponse.RefreshToken = refreshToken
	tokenResponse.AccessToken = accessToken

	return &tokenResponse, nil
}

// Login User
func (service *AuthService) Login(form dto.UserLoginRequest) (*dto.JWTResponse, error) {
	// Check whether user exists
	count, err := service.repository.Count(form.Email)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.ErrUserNotFound
	}

	// Get user by email
	user, err := service.repository.GetUserByEmail(form.Email)
	if err != nil {
		return nil, err
	}

	// Compare the password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password))
	if err != nil {
		return nil, err
	}

	// Create tokens
	refreshToken, accessToken, err := CreateTokens(user.Id, service.config.JWTSecret)
	if err != nil {
		return nil, err
	}
	if refreshToken == "" || accessToken == "" {
		return nil, errors.ErrInvalidTokens
	}

	now := time.Now()

	// Create session
	token := models.UserSessionModel{}
	token.UserId = user.Id
	token.RefreshTokenHash = refreshToken // need to store the hash instead of jwtRefresh token itself - future enhancement
	token.ExpiresAt = now.Add(7 * 24 * time.Hour)
	token.RevokedAt = nil
	token.CreatedAt = now

	if err := service.repository.InsertUserSession(&token); err != nil {
		return nil, err
	}

	tokenResponse := dto.JWTResponse{}
	tokenResponse.RefreshToken = refreshToken
	tokenResponse.AccessToken = accessToken

	return &tokenResponse, nil
}

// Refresh access token
func (service *AuthService) Refresh(form dto.JWTRequest) (*dto.JWTResponse, error) {
	// Check the expiry time of the refresh token
	refreshToken, err := jwt.Parse(form.RefreshToken, func(t *jwt.Token) (any, error) {
		return service.config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !refreshToken.Valid {
		return nil, errors.ErrExpiredToken
	}
	claims := refreshToken.Claims
	
	// Get existing user session
	userSession, err := service.repository.GetUserSessionByRefreshToken(form.RefreshToken)
	if err != nil {
		return nil, err
	}

	// Check whether revoked
	if userSession.RevokedAt != nil {
		return nil, errors.ErrExpiredToken
	}

	if userSession.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, errors.ErrExpiredToken
	}

	// Create tokens
	id, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}
	if id == "" {
		return nil, errors.ErrInvalidCredentials
	}

	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, accessToken, err := CreateTokens(objId, service.config.JWTSecret)
	if err != nil {
		return nil, err
	}
	if accessToken == "" {
		return nil, errors.ErrInvalidTokens
	}

	tokenResponse := dto.JWTResponse{}
	// tokenResponse.RefreshToken = refreshToken
	tokenResponse.AccessToken = accessToken

	return &tokenResponse, nil
}

// Logout user
func (service *AuthService) Logout(form dto.JWTRequest) error {
	// Check the expiry time of the refresh token
	refreshToken, err := jwt.Parse(form.RefreshToken, func(t *jwt.Token) (any, error) {
		return service.config.JWTSecret, nil
	})
	if err != nil {
		return err
	}
	if !refreshToken.Valid {
		return errors.ErrExpiredToken
	}
	
	// Update existing userSession
	err = service.repository.UpdateUserSession(form.RefreshToken, bson.M{"revokedAt": time.Now()})
	
	return err
}

//-------------------------- Reusable functions -----------------------//

func CreateTokens(objId bson.ObjectID, jwtSecret string) (string, string, error) {
	now := time.Now()

	// Refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": objId.Hex(),
		"iat": now.Unix(),
		"exp": now.Add(7 * 24 * time.Hour).Unix(),
	})

	jwtRefresh, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err // user created but unable to create refresh token
	}

	// Access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": objId.Hex(),
		"iat": now.Unix(),
		"exp": now.Add(15 * time.Minute).Unix(),
	})

	jwtAccess, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return jwtRefresh, "", err // user created but unable to create access token
	}

	return jwtRefresh, jwtAccess, nil
}