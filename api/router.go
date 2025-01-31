package api

import (
	"golang_ot/api/v1/healthCheck"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api/v1")
	healthCheck.SetupRoutes(apiGroup)
}
