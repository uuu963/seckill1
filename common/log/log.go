package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog() error {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// 检测是否有目录
	_, err := os.Stat("./logs/")
	if err != nil && os.IsNotExist(err) {
		err := os.Mkdir("./logs/", os.ModePerm)
		if err != nil {
			fmt.Println("无法创建目录", err)
			return err
		}
		fmt.Println("成功创建目录")
	} else if err != nil {
		fmt.Println("出错:", err)
		return err
	} else {
		fmt.Println("目录已存在")
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stdout", "./logs/log.txt",
		},
	}
	logger := zap.Must(config.Build())
	zap.ReplaceGlobals(logger)
	return err
}
