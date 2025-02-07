package user

import (
	"golang_ot/models"
	"golang_ot/models/mongoModels"
	"golang_ot/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{service: services.NewUserService()}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user mongoModels.User
	if err := c.BodyParser(&user); err != nil {
		return models.JSONResponse(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "invalid_request", nil, err.Error())
	}
	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		return models.JSONResponse(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "error_with", []string{"cUser"}, err.Error())
	}
	userResponse := mongoModels.ConvertUserToResponse(createdUser)
	return models.JSONResponse(c, fiber.StatusCreated, fiber.StatusCreated, "success", []string{"cUser"}, userResponse)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		return models.JSONResponse(c, fiber.StatusNotFound, fiber.StatusNotFound, "not_exist", []string{"cUser"}, err.Error())
	}
	userResponse := mongoModels.ConvertUserToResponse(user)
	return models.JSONResponse(c, fiber.StatusOK, fiber.StatusOK, "success", []string{"cUser"}, userResponse)
}
