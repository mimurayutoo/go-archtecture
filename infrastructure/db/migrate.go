package db

import (
	userModel "practice-api/internal/user/model"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	// まずUserテーブルを作成
	if err := DB.AutoMigrate(&userModel.User{}); err != nil {
		panic("failed to migrate user table: " + err.Error())
	}

}
