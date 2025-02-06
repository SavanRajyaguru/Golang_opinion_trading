package user

import (
	m "golang_ot/middleware"
	v "golang_ot/middleware/validators"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	handler := NewUserHandler()
	router.Post("/user/register/v1", m.RequestValidation(&v.RegisterUser{}), handler.CreateUser)
	router.Get("/user/:id/v1", handler.GetUserByID)
}
