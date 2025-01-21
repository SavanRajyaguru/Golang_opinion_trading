package models

import "github.com/gofiber/fiber/v2"

// Response represents the global response structure
type Response struct {
	Status  int         `json:"status"`  // status code
	Message string      `json:"message"` // description of the response
	Data    interface{} `json:"data"`    // additional data
}

// Json response is the utility function to return json response
func JSONResponse(c *fiber.Ctx, statusCode int, jsonStatusCode int, status string, message string, data interface{}) error {
	response := Response{
		Status:  jsonStatusCode,
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}
