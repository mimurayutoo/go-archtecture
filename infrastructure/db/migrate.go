package db

import (
	"practice-api/internal/domain"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&domain.User{}, &domain.Todo{})
}
