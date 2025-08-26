package api

import (
	"E-Commerce-Golang/configs"
	"E-Commerce-Golang/internal/api/rest"
	"E-Commerce-Golang/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v3"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New() // Insatce of fiber
	restHandler := &rest.RestHandler{
		App: app,
	}
	setupRoutes(restHandler)
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)

	//transactionHandler

	// catalogHandler
}
