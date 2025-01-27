package controllers

import (
	"log"

	"github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TestLogin(c *fiber.Ctx) error {
	var loginUser models.Users
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Query Username and Password by loginUser.Username
	var users []models.Users
	result := config.DB.Where("username = ?", loginUser.Username).Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Check User in Valid
	if len(users) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	// Check password
	if bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(loginUser.Password)) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	return c.JSON(fiber.Map{"message": "Login successful", "users": users, "user": loginUser})
}

func Register(c *fiber.Ctx) error {
	var newUser models.Users
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	
	if newUser.Username == "" || newUser.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and Password is required",
		})
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	newUser.Password = string(hashedPassword)

	// Create User
	result := config.DB.Create(&newUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "User created successfully", "user": newUser})
}
