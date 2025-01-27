package controllers

import (
	"github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers ดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.Users

	// Query ข้อมูลผู้ใช้ทั้งหมด
	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{"data": users})
}
