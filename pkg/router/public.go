package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/controller"
)

func PublicRoutes(app *fiber.App) {
	// Create a new group for public routes
	auth := app.Group("/api/v1/auth")

	// Authentication routes
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Delete("/logout", controller.Logout)
	auth.Post("/refresh", controller.RefreshToken)
}
