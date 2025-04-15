package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"size:50;not null;unique"`
	Email        string    `json:"email" gorm:"size:100;not null;unique"`
	PasswordHash string    `json:"-" gorm:"size:255;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

// BeforeSave is a GORM hook that gets called before creating or updating a user
func (u *User) BeforeSave(*gorm.DB) error {
	now := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	return nil
}

// Validate validates the user data
func (u *User) Validate() error {
	// Validate username
	if strings.TrimSpace(u.Username) == "" {
		return ErrInvalidInput("username cannot be empty")
	}
	if len(u.Username) < 3 || len(u.Username) > 50 {
		return ErrInvalidInput("username must be between 3 and 50 characters")
	}

	// Validate email
	if strings.TrimSpace(u.Email) == "" {
		return ErrInvalidInput("email cannot be empty")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return ErrInvalidInput("invalid email format")
	}

	// Validate password hash
	if strings.TrimSpace(u.PasswordHash) == "" {
		return ErrInvalidInput("password cannot be empty")
	}

	return nil
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword checks if the provided password matches the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
