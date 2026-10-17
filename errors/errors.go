package errors

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrDuplicateUserFound = errors.New("user already exists")
	ErrInvalidTokens      = errors.New("Unable to create tokens")
	ErrExpiredToken       = errors.New("Unauthorized: Token has expired")
)
