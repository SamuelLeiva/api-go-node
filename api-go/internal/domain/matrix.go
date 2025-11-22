package domain

type MatrixPayload struct {
	Matrix [][]float64 `json:"matrix"`
}

type QRResult struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}
