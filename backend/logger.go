package backend

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger 创建日志记录对象
func NewLogger(config *LogConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   config.Filename,
		MaxAge:     config.MaxAge,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		Compress:   config.Compress,
		LocalTime:  config.LocalTime,
	}
}
