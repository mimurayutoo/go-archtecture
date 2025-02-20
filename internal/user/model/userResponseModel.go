package userModel

import (
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
	Sex       string         `json:"sex"`
}

func (u *ResponseUser) NewUserResponse(user *User) *ResponseUser {
	u.SetId(user.ID)
	u.SetCreatedAt(user.CreatedAt)
	u.SetUpdatedAt(user.UpdatedAt)
	u.SetDeletedAt(user.DeletedAt)
	u.SetUsername(user.Username)
	u.SetEmail(user.Email)
	u.SetPassword(user.Password)
	u.SetAge(user.Age)
	u.SetSex(user.Sex)
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

func (u *ResponseUser) GetSex() string {
	return u.Sex
}

func (u *ResponseUser) SetSex(sex string) {
	u.Sex = sex
}
