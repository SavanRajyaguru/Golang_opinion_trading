package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/yudiz-savan-rajyaguru/golang-ot/config"
)

func ConnectMySQL(cnf config.MySQLConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to load MySQL database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}
	return db
}
