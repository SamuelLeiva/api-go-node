package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"api-go/internal/domain"
	"api-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct{}

func (h *MatrixHandler) HandleQR(c *fiber.Ctx) error {
	var payload domain.MatrixPayload

	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	q, r, err := services.ProcessQR(payload.Matrix)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result := domain.QRResult{Q: q, R: r}

	// Forward automático a API-Node si se configuró
	nodeURL := os.Getenv("NODE_API_URL")

	if nodeURL != "" {
		body := map[string]any{
			"matrices": []any{result.Q, result.R},
		}

		payloadBytes, _ := json.Marshal(body)

		// obtener token del frontend
		authHeader := c.Get("Authorization")

		// construimos el request con header Authorization
		req, _ := http.NewRequest("POST", nodeURL, bytes.NewReader(payloadBytes))
		req.Header.Set("Content-Type", "application/json")
		if authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Println("❌ Error forwarding to Node API:", err)
			// Devolvemos QR normal, porque no es un error fatal
			return c.JSON(result)
		}

		// Asegurarse de cerrar el body de la respuesta
		defer resp.Body.Close()

		// Parsear respuesta del API-Node
		var stats map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
			log.Println("❌ Error parsing Node response:", err)
			return c.JSON(result)
		}

		log.Println("📨 Forwarded to Node API:", resp.Status)

		return c.JSON(map[string]any{
			"Q":     result.Q,
			"R":     result.R,
			"stats": stats,
		})
	}

	// Si no hay API-Node configurada, respondemos solo con el QR
	return c.JSON(result)
}
