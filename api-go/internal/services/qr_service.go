package services

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

var ErrInvalidMatrix = errors.New("Matriz inválida: debe ser rectangular y no vacía")

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

	// aplanamos la matriz a una dimensión (tipo, n° de elementos referenciados, capacidad)
	flat := make([]float64, 0, rowCount*colCount)
	// por cada fila, agregar sus elementos a la matriz plana
	// usamos ... para desempaquetar la fila
	for _, row := range m {
		flat = append(flat, row...)
	}

	// gonum necesita una matriz plana para crear una matriz densa
	A := mat.NewDense(rowCount, colCount, flat)

	// generamos las matrices Q(ortogonal) y R(triangular superior)
	// a partir de A, las guarda internamente en qr
	var qr mat.QR
	qr.Factorize(A)

	// extraemos Q y R de qr
	var Q, R mat.Dense
	qr.QTo(&Q)
	qr.RTo(&R)

	// convertimos las matrices densas a slices 2D [][]float64
	Qs := denseToSlice(&Q)
	Rs := denseToSlice(&R)

	// limpiar valores extremadamente pequeños
	zeroTiny(Qs)
	zeroTiny(Rs)

	return Qs, Rs, nil
}

// convierte una matriz densa de gonum a un slice 2D [][]float64
func denseToSlice(d *mat.Dense) [][]float64 {
	r, c := d.Dims()
	out := make([][]float64, r)
	for i := range r {
		row := make([]float64, c)
		for j := range c {
			row[j] = d.At(i, j)
		}
		out[i] = row
	}
	return out
}

// reemplaza valores muy pequeños por cero exacto
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
