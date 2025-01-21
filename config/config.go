package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL         MySQLConfig
	Mongo         MongoConfig
	Postgres      PostgresConfig
	DefaultConfig DefaultConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type MongoConfig struct {
	URI      string
	Database string
}

type PostgresConfig struct {
	User            string
	Password        string
	Host            string
	Port            string
	Database        string
	MaxConns        int32
	MinConns        int32
	MaxConnIdleTime time.Duration
}

type DefaultConfig struct {
	Port string
}

// LoadConfig loads configuration using viper
func LoadConfig() Config {
	viper.AutomaticEnv() // Automatically read environment variables

	// Bind environment variables
	viper.BindEnv("mysql.user", "MYSQL_USER")
	viper.BindEnv("mysql.password", "MYSQL_PASSWORD")
	viper.BindEnv("mysql.host", "MYSQL_HOST")
	viper.BindEnv("mysql.port", "MYSQL_PORT")
	viper.BindEnv("mysql.database", "MYSQL_DB")

	viper.BindEnv("mongo.uri", "MONGO_URI")
	viper.BindEnv("mongo.database", "MONGO_DB")

	viper.BindEnv("postgres.user", "POSTGRES_USER")
	viper.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	viper.BindEnv("postgres.host", "POSTGRES_HOST")
	viper.BindEnv("postgres.port", "POSTGRES_PORT")
	viper.BindEnv("postgres.database", "POSTGRES_DB")
	viper.BindEnv("postgres.max_conns", "POSTGRES_MAX_CONNS")
	viper.BindEnv("postgres.min_conns", "POSTGRES_MIN_CONNS")
	viper.BindEnv("postgres.max_conn_idle_time", "POSTGRES_MAX_CONN_IDLE_TIME")

	viper.BindEnv("default_config.port", "PORT") // Ensure correct binding

	// Unmarshal the configuration into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	// Parse duration fields explicitly if needed
	if maxConnIdleTime := viper.GetString("postgres.max_conn_idle_time"); maxConnIdleTime != "" {
		duration, err := time.ParseDuration(maxConnIdleTime)
		if err != nil {
			log.Fatalf("Error parsing max_conn_idle_time: %v", err)
		}
		config.Postgres.MaxConnIdleTime = duration
	}

	return config
}
