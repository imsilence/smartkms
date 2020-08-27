package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/imsilence/smartkms/backend"
	"github.com/imsilence/smartkms/backend/request"
	"github.com/imsilence/smartkms/backend/utils"
)

type dataService struct {
}

// Encrypt 对数据加密
func (s *dataService) Encrypt(req *request.EncryptRequest) (string, error) {
	// 获取应用key密钥
	kek, err := KekService.GetByAppKey(req.AppKey)
	if err != nil {
		return "", err
	}

	// 生成数据密钥
	key, err := utils.RandAesKey()
	if err != nil {
		return "", err
	}

	// 使用数据密钥对数据解密
	dataCiphertext, err := utils.AESGCMEncodeString(key, req.Plaintext)
	if err != nil {
		return "", err
	}

	// 使用根密钥解密key密钥
	keySecret, err := utils.AESGCMDecode(backend.App.Config.Key, kek.AppSecret)
	if err != nil {
		return "", err
	}

	// 使用key密钥对数据密钥进行加密
	keyCiphertext, err := utils.AESGCMEncodeString(string(keySecret), key)
	if err != nil {
		return "", err
	}

	// 将数据密钥密文和数据密文返回
	return fmt.Sprintf("%s$$%s", keyCiphertext, dataCiphertext), nil
}

// Decrypt 对数据解密
func (s *dataService) Decrypt(req *request.DecryptRequest) (string, error) {
	// 分割数据密钥密文和数据密文
	texts := strings.SplitN(req.Ciphertext, "$$", 2)
	if len(texts) < 2 {
		return "", errors.New("ciphertext error")
	}

	keyCiphertext, dataCiphertext := texts[0], texts[1]

	// 获取应用Key密钥
	kek, err := KekService.GetByAppKey(req.AppKey)
	if err != nil {
		return "", err
	}

	// 使用根密钥解密Key密钥
	keySecret, err := utils.AESGCMDecode(backend.App.Config.Key, kek.AppSecret)
	if err != nil {
		return "", err
	}

	// 使用key密钥解密数据密钥密文
	keyPlaintext, err := utils.AESGCMDecode(string(keySecret), keyCiphertext)
	if err != nil {
		return "", err
	}

	// 使用数据密钥解密数据密文
	dataPlaintext, err := utils.AESGCMDecode(string(keyPlaintext), dataCiphertext)
	if err != nil {
		return "", err
	}

	// 返回数据密文
	return string(dataPlaintext), nil
}

// DataService 数据服务对象
var DataService = new(dataService)
