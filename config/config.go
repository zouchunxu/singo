package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("application")
	config.AddConfigPath("./")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
		return
	}
	configs := config.AllSettings()
	for k, v := range configs {
		config.SetDefault(k, v)
	}
	if env != "" {
		config.SetConfigType("yaml")
		config.SetConfigName(fmt.Sprintf("application-%s", env))
		config.AddConfigPath("./")
		err = config.ReadInConfig()
		if err != nil {
			return
		}
	}
}

func GetConfig() *viper.Viper {
	return config
}
