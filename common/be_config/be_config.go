package beconfig

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 获取配置项
func GetConfig(key string) (string, error) {
	logger := zap.L()
	if logger == nil {
		fmt.Println("logger is nil in GetConfig")
		os.Exit(-1)
	}
	if key == "" {
		logger.Error("failed,key is nil")
		return "", fmt.Errorf("failed,key is nil")
	}
	viper.SetConfigFile(".cfg_linux.json")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error("未找到配置文件" + err.Error())
			return "", err
		}
		logger.Error("配置文件错误：" + err.Error())
		return "", err
	}
	// 成功找到文件后看配置项是否存在
	if !viper.IsSet(key) {
		logger.Error("在配置文件中未找到配置项" + key)
		return "", fmt.Errorf("在配置文件中未找到配置项" + key)
	}
	// 读取指定配置项
	value := viper.GetString(key)

	return value, nil
}
