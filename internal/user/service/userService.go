package userService

import (
	"errors"
	"fmt"
	userModel "practice-api/internal/user/model"
	userRepository "practice-api/internal/user/repository"
	userValidator "practice-api/internal/user/validator"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(user *userModel.User) error
}

// NewUserService サービスの初期化
func NewUserService(repo userRepository.IUserRepository) IUserService {
	return &UserService{repository: repo}
}

type UserService struct {
	repository userRepository.IUserRepository
	validator  userValidator.UserValidator
}

func (s UserService) CreateUser(user *userModel.User) error {
	fmt.Println("バリデーション開始")

	if err := s.validator.ValidateUser(*user); err != nil {
		return errors.New("バリデーションに失敗しました")
	}

	hashedPassword, err := (s.hashPassword(user.GetPassword()))

	if err != nil {
		return errors.New("ハッシュ化に失敗しました")
	}
	// 平文からハッシュ化したパスワードに変更
	user.SetPassword(hashedPassword)

	if err := s.repository.Create(*user); err != nil {
		return err
	}
	return nil
}

func (s UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
