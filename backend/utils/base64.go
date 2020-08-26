package utils

import (
	"encoding/base64"
)

// Base64Encode Base64编码
func Base64Encode(data []byte) string {
	return base64.RawStdEncoding.EncodeToString(data)
}

// Base64Decode Base64解码
func Base64Decode(text string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(text)
}
