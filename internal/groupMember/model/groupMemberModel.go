package groupMemberModel

import (
	"time"

	"gorm.io/gorm"
)

type GroupMember struct {
	// GroupIDとUserIDの組み合わせでユニーク制約を設定
	// 同じユーザーが同じグループに複数回所属することを防ぐ
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	GroupID   uint           `gorm:"not null;index"`
	UserID    uint           `gorm:"not null;index"`
	UniqueID  string         `gorm:"uniqueIndex:idx_group_user;type:varchar(255);default:CONCAT(group_id,'-',user_id)"`
}
