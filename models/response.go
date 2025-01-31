package models

import (
	"golang_ot/lang"

	"github.com/gofiber/fiber/v2"
)

// Response represents the global response structure
type Response struct {
	Status  int         `json:"status"`  // status code
	Message string      `json:"message"` // description of the response
	Data    interface{} `json:"data"`    // additional data
}

// Json response is the utility function to return json response
func JSONResponse(c *fiber.Ctx, statusCode int, jsonStatusCode int, message string, replacements []string, data interface{}) error {
	landCode := c.Locals("userLanguage").(string)
	messages := lang.GetMessage(landCode, message)

	// Replace placeholders if needed
	if len(replacements) > 0 {
		messages = lang.ReplacePlaceholders(landCode, messages, replacements...)
	}
	response := Response{
		Status:  jsonStatusCode,
		Message: messages,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}
