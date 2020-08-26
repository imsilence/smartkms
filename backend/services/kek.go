package services

import (
	"github.com/imsilence/smartkms/backend"
	"github.com/imsilence/smartkms/backend/models"
	"github.com/imsilence/smartkms/backend/utils"
)

type kekService struct {
}

// Apply 申请应用密钥
func (s *kekService) Apply() (*models.Kek, error) {
	// 生成应用key密钥
	key, err := utils.RandAesKey()
	if err != nil {
		return nil, err
	}

	// 使用RootKey对key密钥进行加密
	keySecret, err := utils.AESGCMEncode(backend.App.Config.Key, []byte(key))
	if err != nil {
		return nil, err
	}

	// 存储key密钥信息
	kek := &models.Kek{
		AppID:     utils.UUID(),
		AppKey:    utils.UUID(),
		AppSecret: keySecret,
	}

	if err := backend.App.Db.Create(kek).Error; err != nil {
		return nil, err
	}

	return kek, nil
}

// GetByAppKey 获取应用密钥
func (s *kekService) GetByAppKey(appKey string) (*models.Kek, error) {
	kek := &models.Kek{}
	if err := backend.App.Db.Find(kek, "app_key=?", appKey).Error; err != nil {
		return nil, err
	}
	return kek, nil
}

// KekService kek服务对象
var KekService = new(kekService)
