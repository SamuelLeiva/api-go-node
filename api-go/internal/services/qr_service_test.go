package services_test

import (
	"api-go/internal/services"
	"testing"
)

func TestProcessQR_Success(t *testing.T) {
	t.Log("Starting TestProcessQR_Success...")

	matrix := [][]float64{
		{1, 2},
		{3, 4},
	}

	t.Logf("Input matrix: %#v", matrix)

	Q, R, err := services.ProcessQR(matrix)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	t.Logf("Output Q matrix: %#v", Q)
	t.Logf("Output R matrix: %#v", R)

	if len(Q) != 2 || len(R) != 2 {
		t.Fatalf("Expected 2x2 matrices for Q and R")
	}

	t.Log("TestProcessQR_Success completed successfully")
}

func TestProcessQR_InvalidMatrix_Empty(t *testing.T) {
	t.Log("Starting TestProcessQR_InvalidMatrix_Empty...")

	matrix := [][]float64{}

	Q, R, err := services.ProcessQR(matrix)

	t.Logf("Result Q: %#v, R: %#v, err: %v", Q, R, err)

	if err == nil {
		t.Fatal("Expected an error for empty matrix, got nil")
	}

	t.Log("TestProcessQR_InvalidMatrix_Empty OK")
}

func TestProcessQR_InvalidMatrix_NonRectangular(t *testing.T) {
	t.Log("Starting TestProcessQR_InvalidMatrix_NonRectangular...")

	matrix := [][]float64{
		{1, 2, 3},
		{4, 5},
	}

	Q, R, err := services.ProcessQR(matrix)
	t.Logf("Result Q: %#v, R: %#v, err: %v", Q, R, err)

	if err == nil {
		t.Fatal("Expected error for non-rectangular matrix")
	}

	t.Log("TestProcessQR_InvalidMatrix_NonRectangular OK")
}
