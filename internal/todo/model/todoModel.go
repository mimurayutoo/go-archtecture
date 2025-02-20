package todoModel

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `gorm:"type:varchar(255);not null"`
	UserID      uint           `gorm:"not null"` // 外部キー制約を明示的に定義
	Description string         `gorm:"type:text"`
	DueDate     time.Time
	Completed   bool `gorm:"default:false"`
}

// 構造体のゲッター関数とセッター関数を定義してください

// ゲッター関数
func (t *Todo) GetID() uint {
	return t.ID
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) GetUserID() uint {
	return t.UserID
}

func (t *Todo) GetDescription() string {
	return t.Description
}

func (t *Todo) GetDueDate() time.Time {
	return t.DueDate
}

func (t *Todo) GetCompleted() bool {
	return t.Completed
}

// セッター関数
func (t *Todo) SetTitle(title string) {
	t.Title = title
}

func (t *Todo) SetUserID(userID uint) {
	t.UserID = userID
}

func (t *Todo) SetDescription(description string) {
	t.Description = description
}

func (t *Todo) SetDueDate(dueDate time.Time) {
	t.DueDate = dueDate
}

func (t *Todo) SetCompleted(completed bool) {
	t.Completed = completed
}
