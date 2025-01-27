package controllers

import (
	"log"
	"os"
	"strconv"
	"time"

	// "github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(c *fiber.Ctx) error {
	var users models.Users
	if err := c.BodyParser(&users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	mockUser := models.Users{
		ID:        1,
		Username:  "admin",
		Password:  "admin",
		Email:     "",
		CreatedAt: time.Now(),
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	bcryptCostStr := os.Getenv("BCRYPT_COST")
	bcryptCost, err := strconv.Atoi(bcryptCostStr)
	if err != nil {
		log.Fatal("Invalid BCRYPT_COST value")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(mockUser.Password), bcryptCost)
	mockUser.Password = string(hashedPassword)
	
	if users.Username == mockUser.Username && bcrypt.CompareHashAndPassword([]byte(mockUser.Password), []byte(users.Password)) == nil {
		return c.JSON(fiber.Map{"message": "Login successful", "user": mockUser})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}
	// result := config.DB.Find(&users)
	// if result.Error != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": result.Error.Error(),
	// 	})
	// }
}
