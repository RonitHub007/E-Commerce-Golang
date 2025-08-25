package api

import (
	"E-Commerce-Golang/configs"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New() // Insatce of fiber

	app.Get("/health", HealthCheck)
	app.Listen(config.ServerPort)
}

func HealthCheck(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I'm Alive",
	})
}
