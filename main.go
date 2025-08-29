// main.go
package main

import (
	"log"

	"github.com/sjzar/chatlog/cmd/chatlog"
	"github.com/sjzar/chatlog/internal/logger"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 设置全局 logger 写入文件
	logger.SetupGlobalLogger()

	chatlog.Execute()
}
