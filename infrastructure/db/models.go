package db

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Username string `gorm:"type:varchar(255);not null"`
// 	Email    string `gorm:"type:varchar(255);unique;not null"`
// 	Password string `gorm:"type:varchar(255);not null"`
// 	Age      int    `gorm:"type:int"`
// 	Sex      string `gorm:"type:varchar(10)"`
// 	Todos    []Todo `gorm:"foreignKey:UserID"`
// }
// type Todo struct {
// 	gorm.Model
// 	Title       string `gorm:"type:varchar(255);not null"`
// 	UserID      uint   `gorm:"not null"` // 外部キー制約を明示的に定義
// 	Description string `gorm:"type:text"`
// 	DueDate     time.Time
// 	Completed   bool `gorm:"default:false"`
// }
