package healthCheck

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yudiz-savan-rajyaguru/golang-ot/models"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	return models.JSONResponse(c, fiber.StatusOK, fiber.StatusOK, "OK", "Server is running fine", nil)
}
