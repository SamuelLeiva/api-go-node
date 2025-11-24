package main

import (
	"log"
	"os"
	"time"

	"api-go/internal/auth"
	"api-go/internal/handlers"
	"api-go/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // dominio del frontend
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// --- AUTH ---
	authHandler := auth.AuthHandler{}

	matrixHandler := handlers.MatrixHandler{}

	app.Post("/login", authHandler.Login)
	app.Post("/qr", middleware.AuthRequired, matrixHandler.HandleQR)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("API-Go running on port", port)
	log.Fatal(app.Listen(":" + port))
}
