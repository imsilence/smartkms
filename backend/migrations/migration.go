package migrations

import (
	"github.com/imsilence/smartkms/backend/models"
	"github.com/jinzhu/gorm"
)

// Migrate 同步模型到数据库
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Kek{})
}
