package userResponseModel

import (
	userModel "practice-api/internal/user/model"
	"time"

	"gorm.io/gorm"
)

type ResponseUser struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Age       int            `json:"age"`
	Token     string         `json:"token,omitempty"`
}

func (u *ResponseUser) NewUserResponse(user *userModel.User) *ResponseUser {
	u.SetId(user.ID)
	u.SetCreatedAt(user.CreatedAt)
	u.SetUpdatedAt(user.UpdatedAt)
	u.SetDeletedAt(user.DeletedAt)
	u.SetUsername(user.Username)
	u.SetEmail(user.Email)
	u.SetPassword(user.Password)
	u.SetAge(user.Age)
	return u
}

func (u *ResponseUser) NewLoginResponse(user *userModel.User, token string) *ResponseUser {
	u.NewUserResponse(user)
	u.SetToken(token)
	return u
}

func (u *ResponseUser) GetId() uint {
	return u.ID
}

func (u *ResponseUser) SetId(id uint) {
	u.ID = id
}

func (u *ResponseUser) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *ResponseUser) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *ResponseUser) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *ResponseUser) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

func (u *ResponseUser) GetDeletedAt() gorm.DeletedAt {
	return u.DeletedAt
}

func (u *ResponseUser) SetDeletedAt(deletedAt gorm.DeletedAt) {
	u.DeletedAt = deletedAt
}

func (u *ResponseUser) GetUsername() string {
	return u.Username
}

func (u *ResponseUser) GetToken() string {
	return u.Token
}

func (u *ResponseUser) SetUsername(username string) {
	u.Username = username
}

func (u *ResponseUser) GetEmail() string {
	return u.Email
}

func (u *ResponseUser) SetEmail(email string) {
	u.Email = email
}

func (u *ResponseUser) GetPassword() string {
	return u.Password
}

func (u *ResponseUser) SetPassword(password string) {
	u.Password = password
}

func (u *ResponseUser) GetAge() int {
	return u.Age
}

func (u *ResponseUser) SetAge(age int) {
	u.Age = age
}

func (u *ResponseUser) SetToken(token string) {
	u.Token = token
}
