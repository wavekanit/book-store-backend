package main

import (
	"log"

	"book-store-backend/src/config"
	"book-store-backend/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDatabase()

	// สร้างแอป Fiber
	app := fiber.New()

	// ใช้ Logger Middleware
	app.Use(logger.New())

	// ตั้งค่า Routes
	routes.SetupRoutes(app)

	// เริ่มเซิร์ฟเวอร์
	log.Fatal(app.Listen(":3000"))
}
