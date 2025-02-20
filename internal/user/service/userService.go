package userService

import (
	"errors"
	userModel "practice-api/internal/user/model"
	userRepository "practice-api/internal/user/repository"
	userValidator "practice-api/internal/user/validator"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(userInput *userModel.UserInput) error
	GetUserByID(id uint) (*userModel.ResponseUser, error)
}

// NewUserService サービスの初期化
func NewUserService(repo userRepository.IUserRepository) IUserService {
	return &UserService{repository: repo}
}

type UserService struct {
	repository userRepository.IUserRepository
	validator  userValidator.UserValidator
}

func (s UserService) CreateUser(userInput *userModel.UserInput) error {
	// userインスタンスを作成
	user := userModel.User{}
	user.NewUser(userInput)

	// バリデーションを行う
	if err := s.validator.ValidateUser(user); err != nil {
		return errors.New("バリデーションに失敗しました")
	}

	existingUser, _ := s.repository.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("すでに同じユーザーが存在しています。")
	}

	hashedPassword, err := (s.hashPassword(user.GetPassword()))

	if err != nil {
		return errors.New("ハッシュ化に失敗しました")
	}
	// 平文からハッシュ化したパスワードに変更
	user.SetPassword(hashedPassword)

	if err := s.repository.Create(user); err != nil {
		return err
	}
	return nil
}

func (s UserService) GetUserByID(id uint) (*userModel.ResponseUser, error) {
	dbUser, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// DBモデルからドメインモデルへの変換
	var existingUser userModel.ResponseUser
	existingUser.NewUserResponse(dbUser)

	return &existingUser, nil
}

func (s UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
