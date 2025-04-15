package models

import (
	"errors"
	"fmt"
)

// Common errors
var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
)

// ErrInvalidInput returns a formatted error for invalid input
func ErrInvalidInput(msg string) error {
	return fmt.Errorf("invalid input: %s", msg)
}

// ErrDatabase returns a formatted error for database operations
func ErrDatabase(err error) error {
	return fmt.Errorf("database error: %w", err)
}
