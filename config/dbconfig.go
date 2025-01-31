package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type MongoConfig struct {
	URL      string
	Database string
}

type PostgresConfig struct {
	User            string
	Password        string
	Host            string
	Port            string
	Database        string
	MaxConns        int
	MinConns        int
	MaxConnIdleTime time.Duration
}

type DBConfig struct {
	MySQL    MySQLConfig
	Mongo    MongoConfig
	Postgres PostgresConfig
}

// LoadDbConfig loads database configuration using Viper
func LoadDbConfig() DBConfig {
	// Set default values for MySQL
	viper.SetDefault("MYSQL_USER", "root")
	viper.SetDefault("MYSQL_PASSWORD", "root")
	viper.SetDefault("MYSQL_HOST", "127.0.0.1")
	viper.SetDefault("MYSQL_PORT", "3306")
	viper.SetDefault("MYSQL_DB", "test")

	// Set default values for MongoDB
	viper.SetDefault("MONGO_URL", "mongodb://localhost:27017")
	viper.SetDefault("MONGO_DB", "test")

	// Set default values for PostgreSQL
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "password")
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_DB", "test")
	viper.SetDefault("POSTGRES_MAX_CONNS", 10)
	viper.SetDefault("POSTGRES_MIN_CONNS", 2)
	viper.SetDefault("POSTGRES_MAX_CONN_IDLE_TIME", "5m")

	// Bind environment variables
	viper.BindEnv("MYSQL_USER")
	viper.BindEnv("MYSQL_PASSWORD")
	viper.BindEnv("MYSQL_HOST")
	viper.BindEnv("MYSQL_PORT")
	viper.BindEnv("MYSQL_DB")

	viper.BindEnv("MONGO_URL")
	viper.BindEnv("MONGO_DB")

	viper.BindEnv("POSTGRES_USER")
	viper.BindEnv("POSTGRES_PASSWORD")
	viper.BindEnv("POSTGRES_HOST")
	viper.BindEnv("POSTGRES_PORT")
	viper.BindEnv("POSTGRES_DB")
	viper.BindEnv("POSTGRES_MAX_CONNS")
	viper.BindEnv("POSTGRES_MIN_CONNS")
	viper.BindEnv("POSTGRES_MAX_CONN_IDLE_TIME")

	// Unmarshal the configuration into the Config struct
	var dbConfig DBConfig
	if err := viper.Unmarshal(&dbConfig); err != nil {
		log.Fatalf("Error unmarshalling DBConfig: %v", err)
	}

	// Parse duration fields explicitly if needed
	if maxConnIdleTime := viper.GetString("POSTGRES_MAX_CONN_IDLE_TIME"); maxConnIdleTime != "" {
		duration, err := time.ParseDuration(maxConnIdleTime)
		if err != nil {
			log.Fatalf("Error parsing POSTGRES_MAX_CONN_IDLE_TIME: %v", err)
		}
		dbConfig.Postgres.MaxConnIdleTime = duration
	}
	dbConfig.MySQL = MySQLConfig{
		User:     viper.GetString("MYSQL_USER"),
		Password: viper.GetString("MYSQL_PASSWORD"),
		Host:     viper.GetString("MYSQL_HOST"),
		Port:     viper.GetString("MYSQL_PORT"),
		Database: viper.GetString("MYSQL_DB"),
	}
	dbConfig.Mongo = MongoConfig{
		URL:      viper.GetString("MONGO_URL"),
		Database: viper.GetString("MONGO_DB"),
	}
	// Debug output
	// fmt.Printf("Loaded Database Config: %+v\n", dbConfig)

	return dbConfig
}
