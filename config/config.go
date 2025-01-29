package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBconfig         DBConfig
	ThirdPartyConfig ThirdPartyConfig
	DefaultConfig    DefaultConfig
}

func LoadConfig() Config {
	viper.SetConfigFile("../../.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: No .env file found, using system environment variables\n")
	}
	viper.AutomaticEnv() // Automatically read environment variables
	return Config{
		DBconfig:      LoadDbConfig(),
		DefaultConfig: LoadDefaultConfig(),
	}
}
