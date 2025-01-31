package main

import (
	"log"

	"golang_ot/api"
	"golang_ot/config"
	"golang_ot/internal/database"
	"golang_ot/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()
	// log.Println("Config: ", cfg.DBconfig.MySQL)
	// Connect to MySQL
	db := database.ConnectMySQL(cfg.DBconfig.MySQL)
	defer db.Close()
	// log.Println("Connected to MySQL successfully", db)

	// // Connect to MongoDB
	mongoDB := database.ConnectMongo(cfg.DBconfig.Mongo)
	defer mongoDB.Client().Disconnect(nil)
	// log.Println("Connected to MongoDB successfully", mongoDB)

	// // Connect to PostgreSQL
	// pgDB := database.ConnectPostgres(cfg.Postgres)
	// defer pgDB.Close()
	// log.Println("Connected to PostgreSQL successfully", pgDB)

	// Initialize fiber app
	app := fiber.New()

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Change this to restrict origins (e.g., "http://example.com")
		AllowHeaders: "Origin, Content-Type, Accept, Language",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Use(middleware.LanguageMiddleware)

	// mount routes
	api.SetupRoutes(app)

	log.Printf("Server is running at %s", cfg.DefaultConfig.Port)
	if err := app.Listen(":" + cfg.DefaultConfig.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
