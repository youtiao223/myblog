package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type serverConfig struct {
	Mode     string
	HttpPort string
}

type dbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

// 全局配置变量
var (
	ServerConfig serverConfig
	DbConfig     dbConfig
)

// Init 使用application.yml 初始化全局配置
func Init() {
	logrus.Debug("config start init")

	var config = viper.New()
	config.SetConfigName("application")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	// 读取配置文件
	err := config.ReadInConfig()
	if err != nil {
		logrus.Error("can not read config file")
	}

	setConfig(config)

	// 自动热更新配置
	config.WatchConfig()
}

func setConfig(v *viper.Viper) {
	err1 := v.Sub("server").Unmarshal(&ServerConfig)
	if err1 != nil {
		logrus.Error("setConfig error")
	}
	err2 := v.Sub("db").Unmarshal(&DbConfig)
	if err2 != nil {
		logrus.Error("setConfig error")
	}

}
