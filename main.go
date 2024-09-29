package main

import (
	"fmt"
	"os"
	"sec-kill/common/log"
	"sec-kill/user_serveice/db"

	"go.uber.org/zap"
)

func main() {
	// 初始化logger
	err := log.InitLog()
	if err != nil {
		fmt.Println("初始化logger失败", err)
		os.Exit(-1)
	}

	logger := zap.L()
	logger.Info("成功初始化logger")

	// 数据库初始化
	err = db.Init()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// 监听服务
}
