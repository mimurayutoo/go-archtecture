package userModel

import (
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
	Sex       string         `gorm:"type:varchar(10)"`
}

func (u *User) NewUser(userInput *UserInput) *User {
	u.SetUsername(userInput.Username)
	u.SetEmail(userInput.Email)
	u.SetPassword(userInput.Password)
	u.SetAge(userInput.Age)
	u.SetSex(userInput.Sex)
	return u
}

// ゲッター関数
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

func (u *User) GetSex() string {
	return u.Sex
}

// セッター関数
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

func (u *User) SetSex(sex string) {
	u.Sex = sex
}
