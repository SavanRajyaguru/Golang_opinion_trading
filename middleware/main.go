package middleware

import (
	"golang_ot/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// LanguageMiddleware sets the user language based on the request header
func LanguageMiddleware(c *fiber.Ctx) error {
	lang := c.Get("Language") // Get the 'Language' header from the request
	switch lang {
	case "hi":
		c.Locals("userLanguage", "Hindi")
	case "en-us":
		c.Locals("userLanguage", "English")
	default:
		c.Locals("userLanguage", "English") // Default language
	}
	return c.Next() // Proceed to the next handler
}

func RequestValidation(reqBody interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		validate := validator.New()
		// Bind and validate request body
		if err := c.BodyParser(reqBody); err != nil {
			return models.JSONResponse(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "invalid_request", nil, err.Error())
		}

		if err := validate.Struct(reqBody); err != nil {
			return models.JSONResponse(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "invalid_request", nil, err.Error())
		}

		// Validate request params
		params := c.AllParams()
		for key, value := range params {
			if value == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Missing required URL parameter: " + key,
				})
			}
		}

		// Validate query parameters
		queryParams := c.Queries()
		for key, value := range queryParams {
			if value == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Missing required query parameter: " + key,
				})
			}
		}

		// Continue to next middleware/handler
		return c.Next()
	}
}
