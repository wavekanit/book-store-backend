package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wavekanit/book-store-backend/src/models"
	"github.com/wavekanit/book-store-backend/src/services"
)
type AuthController struct {
    authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var loginUser models.Users
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	token, err := ac.authService.Login(loginUser.Username, loginUser.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Login successful", "token": token})
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var newUser models.Users
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := ac.authService.Register(&newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}
