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
		b, _ := json.Marshal(body)

		resp, err := http.Post(nodeURL, "application/json", bytes.NewReader(b))

		if err != nil {
			log.Println("Error forwarding to Node API:", err)
		} else {
			log.Println("Forwarded to Node API:", resp.Status)
		}
	}

	return c.JSON(result)
}
