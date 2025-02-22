package userService

import (
	"errors"
	"fmt"
	"os"
	userModel "practice-api/internal/user/model"
	userRepository "practice-api/internal/user/repository"
	userValidator "practice-api/internal/user/validator"
	userInputModel "practice-api/shared/dto/userDTO/userInput"
	userResponseModel "practice-api/shared/dto/userDTO/userResponse"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(userInput *userInputModel.UserSignUpInput) error
	GetUserByID(id uint) (*userResponseModel.ResponseUser, error)
	Login(loginUser userInputModel.UserLoginInput) (*userResponseModel.ResponseUser, error)
}

// NewUserService サービスの初期化
func NewUserService(repo userRepository.IUserRepository) IUserService {
	return &UserService{repository: repo}
}

type UserService struct {
	repository userRepository.IUserRepository
	validator  userValidator.UserValidator
}

func (s UserService) CreateUser(userInput *userInputModel.UserSignUpInput) error {
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

func (s UserService) GetUserByID(id uint) (*userResponseModel.ResponseUser, error) {
	dbUser, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// DBモデルからドメインモデルへの変換
	var existingUser userResponseModel.ResponseUser
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

func (s UserService) Login(loginUser userInputModel.UserLoginInput) (*userResponseModel.ResponseUser, error) {
	user, err := s.repository.FindByEmail(loginUser.Email)
	if err != nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(loginUser.Password)); err != nil {
		return nil, errors.New("ユーザー認証に失敗しました。")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	var responseLoginUser userResponseModel.ResponseUser
	responseLoginUser.NewLoginResponse(user, token)

	return &responseLoginUser, nil
}

func (s UserService) generateToken(user *userModel.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEYが取得できません。")
	}

	// トークンの有効期限を3時間に設定
	expirationTime := time.Now().Add(3 * time.Hour)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   fmt.Sprintf("%d %s", user.GetId(), user.GetEmail()),
	}

	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// トークンをシークレットキーで署名
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
