package config

import (
	"log"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
	BaseUrl     string `mapstructure:"BASE_URL"`
}

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("app")
	config.SetConfigType("env")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		log.Printf("err read config: %v", err)
	}

	err = config.Unmarshal(&ViperConfig{})
	if err != nil {
		log.Printf("err unmarshal config: %v", err)
	}

	return config
}
