package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yudiz-savan-rajyaguru/golang-ot/config"
)

func ConnectPostgres(cfg config.PostgresConfig) *pgxpool.Pool {
	// Build the DSN
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// Configure connection pool
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse PostgreSQL DSN: %v", err)
	}
	config.MaxConns = int32(cfg.MaxConns)
	config.MinConns = int32(cfg.MinConns)
	config.MaxConnIdleTime = cfg.MaxConnIdleTime

	// Create connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return pool
}
