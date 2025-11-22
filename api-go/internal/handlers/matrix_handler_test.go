package handlers_test

import (
	"api-go/internal/handlers"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	app := fiber.New()
	handler := handlers.MatrixHandler{}
	app.Post("/qr", handler.HandleQR)
	return app
}

func TestHandleQR_Success(t *testing.T) {
	t.Log("Starting TestHandleQR_Success...")

	app := setupApp()

	payload := map[string]any{
		"matrix": [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	body, _ := json.Marshal(payload)
	t.Logf("Sending request payload: %s", string(body))

	req := httptest.NewRequest("POST", "/qr", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}

	t.Logf("HTTP status: %d", resp.StatusCode)

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]any
	json.NewDecoder(resp.Body).Decode(&result)

	t.Logf("Handler response: %#v", result)

	if result["q"] == nil || result["r"] == nil {
		t.Fatal("Expected Q and R in handler response")
	}

	t.Log("TestHandleQR_Success completed")
}

func TestHandleQR_InvalidJSON(t *testing.T) {
	t.Log("Starting TestHandleQR_InvalidJSON...")

	app := setupApp()

	req := httptest.NewRequest("POST", "/qr", bytes.NewBuffer([]byte("INVALID")))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	t.Logf("HTTP status: %d", resp.StatusCode)

	if resp.StatusCode != 400 {
		t.Fatalf("Expected status 400, got %d", resp.StatusCode)
	}

	t.Log("TestHandleQR_InvalidJSON OK")
}

func TestHandleQR_InvalidMatrix(t *testing.T) {
	t.Log("Starting TestHandleQR_InvalidMatrix...")

	app := setupApp()

	payload := map[string]any{
		"matrix": [][]float64{
			{1, 2},
			{3},
		},
	}

	body, _ := json.Marshal(payload)
	t.Logf("Sending payload: %s", string(body))

	req := httptest.NewRequest("POST", "/qr", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	t.Logf("HTTP status: %d", resp.StatusCode)

	if resp.StatusCode != 400 {
		t.Fatalf("Expected 400 for invalid matrix")
	}

	t.Log("TestHandleQR_InvalidMatrix OK")
}

func TestHandleQR_ForwardToNode(t *testing.T) {
	t.Log("Starting TestHandleQR_ForwardToNode...")

	// Fake Node API URL (assuming Node API isn’t running)
	os.Setenv("NODE_API_URL", "http://localhost:9999/stats")

	app := setupApp()

	payload := map[string]any{
		"matrix": [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	body, _ := json.Marshal(payload)
	t.Logf("Sending payload: %s", string(body))

	req := httptest.NewRequest("POST", "/qr", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	t.Logf("HTTP status: %d", resp.StatusCode)

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200 OK even if Node API fails, got %d", resp.StatusCode)
	}

	t.Log("TestHandleQR_ForwardToNode OK")
}
