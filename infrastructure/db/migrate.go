package db

import (
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	// まずUserテーブルを作成
	if err := DB.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate user table: " + err.Error())
	}

	// 次にTodoテーブルを作成（User テーブルへの外部キー制約あり）
	if err := DB.AutoMigrate(&Todo{}); err != nil {
		panic("failed to migrate todo table: " + err.Error())
	}
}
