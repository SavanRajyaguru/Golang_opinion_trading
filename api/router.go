package api

import (
	"golang_ot/api/v1/healthCheck"
	"golang_ot/api/v1/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	healthCheck.SetupRoutes(apiGroup)
	user.SetupUserRoutes(apiGroup)
}
