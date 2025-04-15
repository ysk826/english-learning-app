package models

import (
	"time"

	"gorm.io/gorm"
)

// Sentence represents a sentence created by a user
type Sentence struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
}

// BeforeSave is a GORM hook that gets called before creating or updating a sentence
func (s *Sentence) BeforeSave(*gorm.DB) error {
	now := time.Now()
	if s.CreatedAt.IsZero() {
		s.CreatedAt = now
	}
	s.UpdatedAt = now
	return nil
}

// Validate validates the sentence data
func (s *Sentence) Validate() error {
	if s.Content == "" {
		return ErrInvalidInput("content cannot be empty")
	}
	if s.UserID == 0 {
		return ErrInvalidInput("user ID cannot be empty")
	}
	return nil
}
