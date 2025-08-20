package middleware

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/kevinoctavian/evodka_backend/pkg/config"
)

func JWTProtected() fiber.Handler {
	jwtwareConfig := jwtware.Config{
		SigningKey:    []byte(config.AppCfg().JWTAccessKey),
		ErrorHandler:  jwtError,
		SigningMethod: jwt.SigningMethodHS512.Name,
	}

	return jwtware.New(jwtwareConfig)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"msg": err.Error(),
	})
}
