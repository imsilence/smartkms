package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

const (
	aesKeyLength = 32
	nonceLength  = 12
)

// RandAesKey 随机AES KEY
func RandAesKey() (string, error) {
	key, err := RandKey(aesKeyLength)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// AESGCMEncode 对字节切片AES加密
func AESGCMEncode(key string, plaintext []byte) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceLength)
	if n, err := rand.Read(nonce); err != nil {
		return "", err
	} else if n != nonceLength {
		return "", errors.New("Generate Nonce")
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := aesGcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)
	return Base64Encode(ciphertext), nil
}

// AESGCMEncodeString 对字符串AES加密
func AESGCMEncodeString(key string, text string) (string, error) {
	return AESGCMEncode(key, []byte(text))
}

// AESGCMEncodeB64String 对base64字符串AES加密
func AESGCMEncodeB64String(key string, text string) (string, error) {
	ciphertext, err := Base64Decode(text)
	if err != nil {
		return "", err
	}
	return AESGCMEncode(key, ciphertext)
}

// AESGCMDecode AES解密为字符切片
func AESGCMDecode(key string, text string) ([]byte, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	ciphertext, err := Base64Decode(text)
	if err != nil {
		return nil, err
	}

	pos := len(ciphertext) - nonceLength
	nonce := ciphertext[pos:]

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesGcm.Open(nil, nonce, ciphertext[:pos], nil)
}

// AESGCMDecodeToB64String AES解密为base64字符串
func AESGCMDecodeToB64String(key string, text string) (string, error) {
	plaintext, err := AESGCMDecode(key, text)
	if err != nil {
		return "", err
	}
	return Base64Encode(plaintext), nil
}
