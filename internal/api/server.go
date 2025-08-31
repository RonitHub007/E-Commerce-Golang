package api

import (
	"E-Commerce-Golang/configs"
	"E-Commerce-Golang/internal/api/rest"
	"E-Commerce-Golang/internal/api/rest/handlers"
	"E-Commerce-Golang/internal/domain"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New() // Insatce of fiber
	log.Printf("config DSN %v", config.Dsn)
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Stop executing the application db is not connected properly %v\n", err)
	}
	//run mingration
	db.AutoMigrate(&domain.User{})
	restHandler := &rest.RestHandler{
		App: app,
		DB:  db,
	}
	// connect our database
	setupRoutes(restHandler)
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)

	//transactionHandler

	// catalogHandler
}
