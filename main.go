package main

import (
	"log"

	"github.com/FurkanSamaraz/Golang-Chat/internal/pkg/config"
	"github.com/FurkanSamaraz/Golang-Chat/internal/pkg/controllers"
	"github.com/FurkanSamaraz/Golang-Chat/internal/pkg/middleware"
	"github.com/FurkanSamaraz/Golang-Chat/internal/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	config.InitialiseRedis()
	defer config.InitialiseRedis().Close()

	jwtMiddleware := middleware.NewJWTMiddleware("gizli-anahtar")

	chatService := services.ChatService{DB: config.Connection()}
	chatControllers := controllers.ChatController{Svc: chatService}

	userService := services.UserService{DB: config.Connection()}
	userControllers := controllers.UserController{Svc: userService}

	app.Post("/register", userControllers.RegisterHandler)
	app.Post("/login", userControllers.LoginHandler)

	app.Use(jwtMiddleware.Middleware())

	app.Get("/verify-contact", chatControllers.VerifyContactHandler)
	app.Get("/chat-history", chatControllers.ChatHistoryHandler)
	app.Get("/contact-list", chatControllers.ContactListHandler)

	app.Get("/ws", websocket.New(controllers.WsHandler))

	err := app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}

}
