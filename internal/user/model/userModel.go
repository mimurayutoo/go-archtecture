package userModel

import (
	userInputModel "practice-api/shared/dto/userDTO/userInput"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);unique;not null"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Age       int            `gorm:"type:int"`
}

// ユーザー新規作成の際のバインド
func (u *User) NewUser(userInput *userInputModel.UserSignUpInput) *User {
	u.SetUsername(userInput.Username)
	u.SetEmail(userInput.Email)
	u.SetPassword(userInput.Password)
	u.SetAge(userInput.Age)
	return u
}

// ゲッター関数

func (u *User) GetId() uint {
	return u.ID
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetAge() int {
	return u.Age
}

// セッター関数

func (u *User) SetId(id uint) {
	u.ID = id
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetAge(age int) {
	u.Age = age
}
