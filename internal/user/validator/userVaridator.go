package userValidator

import (
	"errors"
	userModel "practice-api/internal/user/model"
	"regexp"
)

type UserValidator struct{}

func (v UserValidator) ValidateUser(user userModel.User) error {
	if err := v.validateName(user.GetUsername()); err != nil {
		return err
	}
	if err := v.validateEmail(user.GetEmail()); err != nil {
		return err
	}
	if err := v.validatePassword(user.GetPassword()); err != nil {
		return err
	}
	if err := v.validateAge(user.GetAge()); err != nil {
		return err
	}
	return nil
}

func (v UserValidator) validateName(name string) error {
	// 名前は必須で、4文字以上である必要がある。
	print(len(name))
	if name == "" || len(name) <= 4 {
		return errors.New("usernameは必須で、4文字以上である必要があります")
	}
	return nil
}

func (v UserValidator) validateEmail(email string) error {
	// Emailの正規表現
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// 正規表現をコンパイル
	re := regexp.MustCompile(emailPattern)

	// メールアドレスが正しい形式かをチェック
	if !re.MatchString(email) {
		return errors.New("メールアドレスの形式が正しくありません")
	}

	return nil
}

func (v UserValidator) validatePassword(password string) error {
	// パスワードは必須で、8文字以上である必要がある。
	if password == "" || len(password) <= 8 {
		return errors.New("passwordは必須で、8文字以上である必要があります")
	}
	return nil
}

func (v UserValidator) validateAge(age int) error {
	// ageは必須で、0以上である必要がある。
	if age <= 0 {
		return errors.New("ageは必須で、0以上である必要があります")
	}

	return nil
}
