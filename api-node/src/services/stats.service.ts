export interface Stats {
  max: number;
  min: number;
  sum: number;
  average: number;
  isDiagonal: boolean;
}

export function calculateStatistics(matrices: number[][][]): Stats {
  let flat: number[] = [];

  // 1. aplanar todas las matrices en una sola lista
  for (const matrix of matrices) {
    for (const row of matrix) {
      flat.push(...row);
    }
  }

  if (flat.length === 0) {
    throw new Error("Datos de matriz vacíos");
  }

  // cálculo de máximo, minimo, suma y promedio
  const max = Math.max(...flat);
  const min = Math.min(...flat);
  const sum = flat.reduce((acc, v) => acc + v, 0);
  const average = sum / flat.length;

  // verificar si alguna es diagonal
  const isDiagonal = matrices.some(isMatrixDiagonal);

  return {
    max,
    min,
    sum,
    average,
    isDiagonal,
  };
}

function isMatrixDiagonal(matrix: number[][]): boolean {
  const rows = matrix.length;
  const cols = matrix[0].length;

  if (rows !== cols) return false;

  // por cada elemento fuera de la diagonal, verificar si es cero
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (i !== j && matrix[i][j] !== 0) return false;
    }
  }

  return true;
}
