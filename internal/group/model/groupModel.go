package groupModel

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text;not null"`
	OwnerID     uint           `gorm:"not null;index"`
	MemberCount uint           `gorm:"type:int;not null;default:0"`
}
