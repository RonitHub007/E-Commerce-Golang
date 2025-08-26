package handlers

import (
	"E-Commerce-Golang/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	// svc UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App
	handler := &UserHandler{}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.Verify)
	app.Post("/get-verification-code", handler.GetVerificationCode)

	app.Post("/profile", handler.UpdateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)

	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrderByID)

	app.Post("/become-seller", handler.BecomeSeller)
}

// ---------- Handler Functions ----------

func (h *UserHandler) Register(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "register",
	})
}

func (h *UserHandler) Login(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login",
	})
}

func (h *UserHandler) Verify(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "verify",
	})
}

func (h *UserHandler) GetVerificationCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get verification code",
	})
}

func (h *UserHandler) UpdateProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "update profile",
	})
}

func (h *UserHandler) GetProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get profile",
	})
}

func (h *UserHandler) AddToCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) GetCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get cart",
	})
}

func (h *UserHandler) GetOrders(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) GetOrderByID(ctx fiber.Ctx) error {
	orderID := ctx.Params("id") // capture order id from route
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get order by id",
		"orderID": orderID,
	})
}

func (h *UserHandler) BecomeSeller(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "become seller",
	})
}
