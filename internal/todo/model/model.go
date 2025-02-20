package todoModel

import (
	"time"
)

type Todo struct {
	id          uint   `gorm:"primaryKey"`
	title       string `gorm:"type:varchar(255);not null"`
	userID      uint   `gorm:"not null"` // 外部キー制約を明示的に定義
	description string `gorm:"type:text"`
	dueDate     time.Time
	completed   bool `gorm:"default:false"`
}

// 構造体のゲッター関数とセッター関数を定義してください

// ゲッター関数
func (t *Todo) GetID() uint {
	return t.id
}

func (t *Todo) GetTitle() string {
	return t.title
}

func (t *Todo) GetUserID() uint {
	return t.userID
}

func (t *Todo) GetDescription() string {
	return t.description
}

func (t *Todo) GetDueDate() time.Time {
	return t.dueDate
}

func (t *Todo) GetCompleted() bool {
	return t.completed
}

// セッター関数
func (t *Todo) SetTitle(title string) {
	t.title = title
}

func (t *Todo) SetUserID(userID uint) {
	t.userID = userID
}

func (t *Todo) SetDescription(description string) {
	t.description = description
}

func (t *Todo) SetDueDate(dueDate time.Time) {
	t.dueDate = dueDate
}

func (t *Todo) SetCompleted(completed bool) {
	t.completed = completed
}
