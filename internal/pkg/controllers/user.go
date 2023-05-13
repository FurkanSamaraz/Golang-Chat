package controllers

import (
	"fmt"
	api_structure "main/internal/pkg/structures"

	api_model "main/internal/pkg/model"
	api_service "main/internal/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Svc api_service.UserService
}

// RegisterHandler godoc
// @Summary       Register User
// @Description   Registers a new user
// @Tags          User
// @Accept        json
// @Produce       json
// @Param         body body api_structure.User true "Request body"
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /register [post]
func (controller *UserController) RegisterHandler(c *fiber.Ctx) error {
	var user api_structure.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Invalid request",
		})
	}
	res := api_model.Register(&user)
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Username and password are required",
		})
	}

	result, rerr := controller.Svc.Register(user)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: rerr.Error(),
		})
	}

	fmt.Println(result)
	res.Message = result.Message

	return c.Status(fiber.StatusOK).JSON(res)
}

// LoginHandler godoc
// @Summary       Login User
// @Description   Logs in a user
// @Tags          User
// @Accept        json
// @Produce       json
// @Param         body body api_structure.User true "Request body"
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /login [post]
func (controller *UserController) LoginHandler(c *fiber.Ctx) error {
	var user api_structure.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Invalid request",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Username and password are required",
		})
	}

	result, rerr := controller.Svc.Login(user)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Create Data",
			Message: rerr.Error(),
		})
	}
	c.Locals("Authorization", result)

	res := api_model.Login(&user)
	res.Jwt = result

	return c.Status(fiber.StatusOK).JSON(res)
}
