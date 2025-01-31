package database

import (
	"database/sql"
	"fmt"
	"log"

	"golang_ot/config"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectMySQL(cfg config.MySQLConfig) *sql.DB {
	// Corrected DSN format
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	// Open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open MySQL connection: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	log.Println("✅ Successfully connected to MySQL!", cfg.Database)
	return db
}
