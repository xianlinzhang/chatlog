package chatlog

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Debug bool

func initLog(cmd *cobra.Command, args []string) {

	log.Info().Msg("initLog is executing")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// 创建 logs 目录（如果不存在）
	logDir := "logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Error().Msg("Failed to create logs directory")
		return
	}

	// 设置 lumberjack 日志轮转
	logFileName := filepath.Join(logDir, "chatlog.log")
	fileWriter := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    100, // megabytes
		MaxAge:     30,  // days
		MaxBackups: 10,
		Compress:   true,
	}

	// 控制台输出
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

	// 同时输出到控制台和文件
	multiWriter := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	log.Logger = log.Output(multiWriter)

	log.Info().Msg("initLog has executed")
}

func initTuiLog(cmd *cobra.Command, args []string) {

	log.Info().Msg("initTuiLog is executing")

	// 保持原有逻辑不变
	//logOutput := io.Discard
	//
	//debug, _ := cmd.Flags().GetBool("debug")
	//if debug {
	//	logpath := util.DefaultWorkDir("")
	//	util.PrepareDir(logpath)
	//	logFD, err := os.OpenFile(filepath.Join(logpath, "chatlog.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	//	if err != nil {
	//		panic(err)
	//	}
	//	logOutput = logFD
	//}
	//
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: logOutput, NoColor: true, TimeFormat: time.RFC3339})
	//logrus.SetOutput(logOutput)

	log.Info().Msg("initTuiLog has executed")
}
