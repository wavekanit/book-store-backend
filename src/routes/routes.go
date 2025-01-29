package routes

import (
	"github.com/wavekanit/book-store-backend/src/controllers"
	"github.com/wavekanit/book-store-backend/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Post("/api/login", controllers.Login)
	v1.Post("/api/register", controllers.Register)

	protected := v1.Group("/").Use(middlewares.AuthenticateToken)
	protected.Get("/api/users/getAllUsers", controllers.GetAllUsers)
	protected.Get("/api/books/getAllBooks", controllers.GetAllBooks)

	admin := v1.Group("/").Use(middlewares.AuthenticateToken, middlewares.CheckAuthLevel(1))
    admin.Get("/api/books/getAllBooks", controllers.GetAllBooks)

}
