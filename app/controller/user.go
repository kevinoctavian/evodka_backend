package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/app/dto"
	"github.com/kevinoctavian/evodka_backend/app/model"
	"github.com/kevinoctavian/evodka_backend/app/repository"
	"github.com/kevinoctavian/evodka_backend/platform/database"
)

func GetUsers(c *fiber.Ctx) error {
	userRepo := repository.NewUserRepo(database.GetDB())
	users, err := userRepo.All(0, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "can't get all users",
		})
	}

	userLists := []dto.User{}
	for _, v := range users {
		userLists = append(userLists, dto.User{
			PublicID:  v.ID,
			Name:      v.Username,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	// Logic to get all users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   userLists[:],
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	updateUser := &model.UpdateUser{}
	if err := c.BodyParser(updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid request body",
		})
	}
	// Logic to update a user

	userRepo := repository.NewUserRepo(database.GetDB())
	err := userRepo.Update(id, updateUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "failed to update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "User updated successfully",
		"id":  id,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	// Logic to delete a user
	id := c.Params("id")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "User deleted successfully",
		"id":  id,
	})
}
