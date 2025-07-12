package routes

import (
	"github.com/wavekanit/book-store-backend/src/config"
	"github.com/wavekanit/book-store-backend/src/controllers"
	"github.com/wavekanit/book-store-backend/src/middlewares"
	"github.com/wavekanit/book-store-backend/src/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1/api")

	// Create a instance of controllers, middleware and services
	authService := services.NewAuthService(config.DB)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(config.DB)
	userController := controllers.NewUserController(userService)

	authMiddleware := middlewares.NewAuthMiddleware()

    // Public routes
	v1.Post("/login", authController.Login)
	v1.Post("/register", authController.Register)

    // Protected routes
	protected := v1.Group("/").Use(authMiddleware.AuthenticateToken)

    // User routes
	user := protected.Group("/user")
	user.Get("/all", userController.GetAllUsers)
    user.Get("/:id", userController.GetUserByID)
}
