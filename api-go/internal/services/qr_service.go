package services

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

var ErrInvalidMatrix = errors.New("matriz inválida: debe ser rectangular y no vacía")

// funcion para procesar la descomposición QR
func ProcessQR(m [][]float64) ([][]float64, [][]float64, error) {
	if len(m) == 0 || len(m[0]) == 0 {
		return nil, nil, ErrInvalidMatrix
	}

	rowCount := len(m)
	colCount := len(m[0])

	// validar que las filas tengan la misma longitud
	for _, row := range m {
		if len(row) != colCount {
			return nil, nil, ErrInvalidMatrix
		}
	}

	// convertir a matriz densa de gonum
	flat := make([]float64, 0, rowCount*colCount)
	for _, row := range m {
		flat = append(flat, row...)
	}

	A := mat.NewDense(rowCount, colCount, flat)

	var qr mat.QR
	qr.Factorize(A)

	var Q, R mat.Dense
	qr.QTo(&Q)
	qr.RTo(&R)

	Qs := denseToSlice(&Q)
	Rs := denseToSlice(&R)

	zeroTiny(Qs)
	zeroTiny(Rs)

	return Qs, Rs, nil
}

func denseToSlice(d *mat.Dense) [][]float64 {
	r, c := d.Dims()
	out := make([][]float64, r)
	for i := 0; i < r; i++ {
		row := make([]float64, c)
		for j := 0; j < c; j++ {
			row[j] = d.At(i, j)
		}
		out[i] = row
	}
	return out
}

func zeroTiny(m [][]float64) {
	eps := 1e-9
	for i := range m {
		for j := range m[i] {
			if math.Abs(m[i][j]) < eps {
				m[i][j] = 0
			}
		}
	}
}
