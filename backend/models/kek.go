package models

import (
	"github.com/jinzhu/gorm"
)

// Kek 密钥加密密钥
type Kek struct {
	gorm.Model

	AppID     string `gorm:"column:app_id;type:varchar(256);not null;default:'';"`
	AppKey    string `gorm:"column:app_key;type:varchar(256);not null;default:'';"`
	AppSecret string `gorm:"column:app_secret;type:text;not null;"`
}
