package controllers

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Login(c *fiber.Ctx) error {
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

	// Generate JWT
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24-hour expiration
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	result = config.DB.Model(&users[0]).Update("token", tokenString)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update token",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Login successful"})
}

func Register(c *fiber.Ctx) error {
	var newUser models.Users
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Fill All of the Fields",
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

	c.JSON(fiber.Map{"message": "User created successfully"})

	return c.Next()
}
