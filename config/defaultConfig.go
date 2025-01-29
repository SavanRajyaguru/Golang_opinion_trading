package config

import (
	"github.com/spf13/viper"
)

type DefaultConfig struct {
	Port string
}

func LoadDefaultConfig() DefaultConfig {
	viper.SetDefault("PORT", "5001")
	viper.BindEnv("PORT")
	defaultConfig := DefaultConfig{
		Port: viper.GetString("PORT"),
	}
	return defaultConfig
}
