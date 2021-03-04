package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	// 初始化 配置文件

	//设置配置文件的名字
	viper.SetConfigName("config")

	//添加配置文件所在的路径
	viper.AddConfigPath("./")

	//设置配置文件类型，可选
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
}

// Get 获得配置
func Get(key string) (value interface{}) {
	return viper.Get(key)
}
