package models

import "time"

// Model 基础模型
type Model struct {
	ID        uint64     `gorm:"auto_increment;primary_key;column:id;" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;not null;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;" json:"-"`
}
