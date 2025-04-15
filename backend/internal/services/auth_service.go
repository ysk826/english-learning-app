package services

import (
	"english-learning-app/internal/models"
	"english-learning-app/internal/repository"
	"errors"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

// authService implements the AuthService interface
type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new authentication service
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Register registers a new user
func (s *authService) Register(username, email, password string) (*models.User, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Check if username already exists
	existingUser, err = s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// Create new user
	user := &models.User{
		Username: username,
		Email:    email,
	}

	// Set password hash
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	// Save user to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user
func (s *authService) Login(email, password string) (*models.User, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
