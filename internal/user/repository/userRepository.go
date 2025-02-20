package userRepository

import (
	"errors"
	"fmt"
	"practice-api/infrastructure/db"
	userModel "practice-api/internal/user/model"

	"gorm.io/gorm"
)

// NewUserRepository リポジトリの初期化
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

type IUserRepository interface {
	Create(user userModel.User) error
	FindByID(id uint) (*db.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

// ユーザーを作成する
func (r UserRepository) Create(user userModel.User) error {
	fmt.Println("リポジトリ ユーザー作成開始")
	result := r.db.Create(&user)
	if result.Error != nil {
		fmt.Println("ユーザー作成エラー:", result.Error)
		return errors.New("データベースにデータを登録する際にエラーが起きました。")
	}
	fmt.Println("ユーザー作成成功")
	return nil
}

// ユーザーIDをもとにユーザーを取得する
func (r UserRepository) FindByID(id uint) (*db.User, error) {
	var user db.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &db.User{}, nil
}
