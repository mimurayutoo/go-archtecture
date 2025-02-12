package domain

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}
