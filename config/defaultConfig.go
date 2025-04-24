package config

import (
	"github.com/spf13/viper"
)

type DefaultConfig struct {
	Port           string
	EncryptionKey  string
	EncryptionIV   string
	JWTSecretAdmin string
	JWTSecretUser  string
}

func LoadDefaultConfig() DefaultConfig {
	viper.SetDefault("PORT", "5001")
	viper.SetDefault("ENCRYPTION_KEY", "38eec576dbcd116d1af348870f2cddbdd032b8558f73f9770c89bd42b094acff")
	viper.SetDefault("ENCRYPTION_IV", "54189a9decdc5aa0e8e1cf1e579d8651")
	viper.SetDefault("JWT_SECRET_ADMIN", "OtAdmin07")
	viper.SetDefault("JWT_SECRET_USER", "OtUser007")
	viper.BindEnv("PORT")
	defaultConfig := DefaultConfig{
		Port:           viper.GetString("PORT"),
		EncryptionKey:  viper.GetString("ENCRYPTION_KEY"),
		EncryptionIV:   viper.GetString("ENCRYPTION_IV"),
		JWTSecretAdmin: viper.GetString("JWT_SECRET_ADMIN"),
		JWTSecretUser:  viper.GetString("JWT_SECRET_USER"),
	}
	return defaultConfig
}
