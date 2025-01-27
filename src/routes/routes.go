package routes

import (
	"github.com/wavekanit/book-store-backend/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/users", controllers.GetAllUsers)
	app.Post("/api/login", controllers.TestLogin)
}
