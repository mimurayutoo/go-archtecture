package groupRequestModel

import (
	"time"

	"gorm.io/gorm"
)

type GroupRequest struct {
	// 同じユーザーが同じグループに対して複数のリクエストを送れないようにする
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	GroupID   uint           `gorm:"not null;index"`
	UserID    uint           `gorm:"not null;index"`
	Status    string         `gorm:"type:enum('pending','approved','rejected');not null;default:'pending'"`
	UniqueID  string         `gorm:"uniqueIndex:idx_group_request_user;type:varchar(255);default:CONCAT(group_id,'-',user_id)"`
}
