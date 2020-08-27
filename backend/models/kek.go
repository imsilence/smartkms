package models

import (
	"github.com/imsilence/smartkms/backend/common/models"
)

// Kek 密钥加密密钥
type Kek struct {
	models.Model

	AppID     string `gorm:"column:app_id;type:varchar(256);not null;default:'';" json:"app_id"`
	AppKey    string `gorm:"column:app_key;type:varchar(256);not null;default:'';" json:"app_key"`
	AppSecret string `gorm:"column:app_secret;type:text;not null;" json:"-"`
}
