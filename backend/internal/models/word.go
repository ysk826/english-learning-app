package models

import (
	"time"

	"gorm.io/gorm"
)

// Word represents a word saved by a user
type Word struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Word      string    `json:"word" gorm:"size:100;not null"`
	Meaning   string    `json:"meaning" gorm:"type:text"`
	Example   string    `json:"example" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
}

// BeforeSave is a GORM hook that gets called before creating or updating a word
func (w *Word) BeforeSave(*gorm.DB) error {
	now := time.Now()
	if w.CreatedAt.IsZero() {
		w.CreatedAt = now
	}
	w.UpdatedAt = now
	return nil
}

// Validate validates the word data
func (w *Word) Validate() error {
	if w.Word == "" {
		return ErrInvalidInput("word cannot be empty")
	}
	if w.UserID == 0 {
		return ErrInvalidInput("user ID cannot be empty")
	}
	return nil
}
