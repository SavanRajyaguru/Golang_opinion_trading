package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yudiz-savan-rajyaguru/golang-ot/api/v1/healthCheck"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api/v1")
	healthCheck.SetupRoutes(apiGroup)
}
