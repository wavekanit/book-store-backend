package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	JWTSecret []byte
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		JWTSecret: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (am *AuthMiddleware) AuthenticateToken(c *fiber.Ctx) error {
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
		return am.JWTSecret, nil
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

	c.Locals("user", claims)
	return c.Next()
}
