package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AuthRequired(c *fiber.Ctx) error {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ No .env file found, continuing with Docker environment variables")
	}

	secret := os.Getenv("JWT_SECRET")
	//secret := "superSecretKey518485151H8t+ds*&^%$#@!"
	if secret == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "JWT secret not configured")
	}

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenStr := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	return c.Next()
}
