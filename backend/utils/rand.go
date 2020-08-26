package utils

import (
	"crypto/rand"
	"errors"
	mrand "math/rand"
	"strings"
	"time"
)

var letters = `abcdefghijklmnopqrstuvwxyz1234567890-=[];,.~!@#$%^&*()_+{}:<>?ABCDEFHIJKLMNOPQRSTUVWXYZ`

func init() {
	mrand.Seed(time.Now().UnixNano())
}

// RandString 随机字符串
func RandString(length int) string {
	count := len(letters)
	var builder strings.Builder
	for length > 0 {
		builder.WriteByte(letters[mrand.Intn(count)])
		length--
	}
	return builder.String()
}

// RandKey 随机字节
func RandKey(length int) ([]byte, error) {
	key := make([]byte, length)
	if n, err := rand.Read(key); err != nil {
		return nil, err
	} else if n != length {
		return nil, errors.New("Generate Key")
	}
	return key, nil
}
