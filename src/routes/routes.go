package routes

import (
	"book-store-backend/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Route สำหรับดึงข้อมูลผู้ใช้ทั้งหมด
	app.Get("/api/users", handlers.GetAllUsers)
}
