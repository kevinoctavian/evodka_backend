package middleware

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(app *fiber.App) {
	// Register the logger middleware
	app.Use(
		cors.New(),
		logger.New(),
	)
}

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	IsAdmin, ok := claims["role"].(string)
	if !ok || IsAdmin != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"msg": "Forbidden",
		})
	}

	return c.Next()
}
