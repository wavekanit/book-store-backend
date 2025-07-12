package controllers

import (
	"github.com/wavekanit/book-store-backend/src/models"
	"github.com/wavekanit/book-store-backend/src/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	var users []models.Users
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(fiber.Map{"message": "Users retrieved successfully", "data": users})
}

func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{"message": "User retrieved successfully", "data": user})
}
