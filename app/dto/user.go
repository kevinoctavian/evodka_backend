package dto

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/model"
)

type User struct {
	PublicID  string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToJson(model model.User) *fiber.Map {
	return &fiber.Map{
		"id":         model.ID,
		"name":       model.Username,
		"email":      model.Email,
		"created_at": model.CreatedAt,
		"updated_at": model.UpdatedAt,
	}
}
