package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yudiz-savan-rajyaguru/golang-ot/api"
	"github.com/yudiz-savan-rajyaguru/golang-ot/config"
)

func main() {
	cfg := config.LoadConfig()

	// // Connect to MySQL
	// db := database.ConnectMySQL(cfg.MySQL)
	// defer db.Close()
	// log.Println("Connected to MySQL successfully", db)

	// // Connect to MongoDB
	// mongoDB := database.ConnectMongo(cfg.Mongo)
	// // defer mongoDB.Client().Disconnect(nil)
	// log.Println("Connected to MongoDB successfully", mongoDB)

	// // Connect to PostgreSQL
	// pgDB := database.ConnectPostgres(cfg.Postgres)
	// defer pgDB.Close()
	// log.Println("Connected to PostgreSQL successfully", pgDB)

	// Initialize fiber app
	app := fiber.New()

	// mount routes
	api.SetupRoutes(app)

	log.Printf("Server is running at %s", cfg.DefaultConfig.Port)
	if err := app.Listen(":" + cfg.DefaultConfig.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
