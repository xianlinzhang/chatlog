package logger

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// SetupLogger 创建并返回配置好的 logger
func SetupLogger() zerolog.Logger {
	// 创建 logs 目录（如果不存在）
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		// 如果创建目录失败，使用控制台输出并记录错误
		consoleLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		consoleLogger.Error().Err(err).Msg("Failed to create logs directory, using console logging")
		return consoleLogger
	}

	// 使用 lumberjack 实现日志轮转
	logFileName := filepath.Join(logDir, "chatlog.log")
	rotatingLog := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    100, // megabytes
		MaxAge:     30,  // days
		MaxBackups: 10,
		Compress:   true,
	}

	// 创建带轮转功能的 logger
	logger := zerolog.New(rotatingLog).With().Timestamp().Logger()
	return logger
}

// SetupGlobalLogger 配置全局 logger
func SetupGlobalLogger() {
	logger := SetupLogger()
	zerolog.DefaultContextLogger = &logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("Logger initialized") // 这应该会写入到文件中
}
