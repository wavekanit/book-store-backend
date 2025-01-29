package controllers

import (
	"github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/models"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	return c.SendString("Add Book")
}

func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Books

	result := config.DB.Find(&books)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully query books data.",
		"data":    books,
	})
}
