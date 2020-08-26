package utils

import (
	"github.com/google/uuid"
)

// UUID uuid生成器
func UUID() string {
	return uuid.New().String()
}
