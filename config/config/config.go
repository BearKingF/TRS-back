package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config = viper.New()

func Init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.WatchConfig() // 自动将配置读入Config变量

	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not found: ", err)
	}
}
