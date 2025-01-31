package middleware

import "github.com/gofiber/fiber/v2"

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
