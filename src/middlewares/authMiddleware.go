package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func AuthenticateToken(c *fiber.Ctx) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(http.StatusForbidden, "Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Failed to authenticate token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Failed to authenticate token",
		})
	}

	c.Locals("claims", claims)
	c.Locals("auth_level", claims["auth_level"])
	c.Locals("username", claims["username"])
	return c.Next()
}

func CheckAuthLevel(level float64) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authLevel := c.Locals("auth_level").(float64)
        if authLevel < level {
            return c.Status(http.StatusForbidden).JSON(fiber.Map{
                "error": "You have no permission",
            })
        }
        return c.Next()
    }
}
