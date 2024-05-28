package model

import "time"

// gorm.Model 定义
type Model struct {
	ID        uint `gorm:"primary_key"` // 定义主键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
