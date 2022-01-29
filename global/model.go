package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID       uint      `gorm:"primarykey"` // 主键ID
	CreateAt time.Time // 创建时间
	UpdateAt time.Time // 更新是时间
	DeleteAt gorm.DeletedAt
}
