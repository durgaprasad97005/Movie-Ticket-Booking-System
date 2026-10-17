package errors

import "errors"

var (
	ErrInvalidCredentials  = errors.New("Invalid credentials")
	ErrUserNotFound        = errors.New("User not found")
	ErrDuplicateUserFound  = errors.New("User already exists")
	ErrInvalidTokens       = errors.New("Unable to create tokens")
	ErrExpiredToken        = errors.New("Unauthorized: Token has expired")
	ErrUserSessionNotFound = errors.New("User session not found")
)
