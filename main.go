package main

import (
	"log"

	"github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.ConnectDatabase()

	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
