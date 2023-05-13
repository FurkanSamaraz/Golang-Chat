package controllers

import (
	api_service "main/internal/pkg/services"

	api_structure "main/internal/pkg/structures"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	Svc api_service.ChatService
}

// VerifyContactHandler godoc
// @Summary       Verify Contact
// @Description   Verifies the contact of a user
// @Tags          User
// @Security      BearerAuth
// @Accept        json
// @Produce       json
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /verify-contact [get]
func (controller *ChatController) VerifyContactHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	res, err := controller.Svc.VerifyContact(claims["username"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Contact Data",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}

// ChatHistoryHandler godoc
// @Summary       Chat History
// @Description   Retrieves the chat history between two users
// @Tags          Chat
// @Accept        json
// @Produce       json
// @Param         u1 query string true "Username of user 1"
// @Param         u2 query string true "Username of user 2"
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /chat-history [get]
func (controller *ChatController) ChatHistoryHandler(c *fiber.Ctx) error {
	username1 := c.Query("u1")
	username2 := c.Query("u2")

	res, err := controller.Svc.ChatHistory(username1, username2)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Chat History Data",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}

// ContactListHandler godoc
// @Summary       Contact List
// @Description   Retrieves the contact list of a user
// @Tags          User
// @Security      BearerAuth
// @Accept        json
// @Produce       json
// @Success       200 {object} api_structure.Response
// @Failure       500 {object} api_structure.ErrorResponse
// @Router        /contact-list [get]
func (controller *ChatController) ContactListHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	res, err := controller.Svc.ContactList(claims["username"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api_structure.ErrorResponse{
			Type:    "Chat List Data",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}
