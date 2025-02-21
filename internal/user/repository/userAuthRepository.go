package userRepository

import (
	"errors"
	userModel "practice-api/internal/user/model"

	"gorm.io/gorm"
)

// NewUserRepository リポジトリの初期化
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

type IUserRepository interface {
	Create(user userModel.User) error
	FindByID(id uint) (*userModel.User, error)
	FindByEmail(email string) (*userModel.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

// ユーザーを作成する
func (r UserRepository) Create(user userModel.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return errors.New("データベースにデータを登録する際にエラーが起きました。")
	}

	return nil
}

// ユーザーIDをもとにユーザーを取得する
func (r UserRepository) FindByID(id uint) (*userModel.User, error) {
	var user userModel.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("ユーザーが見つかりません")
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r UserRepository) FindByEmail(email string) (*userModel.User, error) {
	var user userModel.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
