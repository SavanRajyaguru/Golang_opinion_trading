package healthCheck

import (
	"golang_ot/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	return models.JSONResponse(c, fiber.StatusOK, fiber.StatusOK, "server_runs", nil, time.Now().Local())
}
