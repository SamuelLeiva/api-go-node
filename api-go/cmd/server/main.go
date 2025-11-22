package main

import (
	"log"
	"os"
	"time"

	"api-go/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	handler := handlers.MatrixHandler{}

	app.Post("/qr", handler.HandleQR)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("API-Go running on port", port)
	log.Fatal(app.Listen(":" + port))
}
