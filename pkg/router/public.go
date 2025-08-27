package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/controller"
)

func PublicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg":    "",
			"status": 200,
		})
	})

	// Create a new group for public routes
	auth := app.Group("/api/v1/auth")

	// Authentication routes
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Delete("/logout", controller.Logout)
	auth.Post("/refresh", controller.RefreshToken)

	// statistic routes
	// statistic := app.Group("/api/v1/statistic")
}
