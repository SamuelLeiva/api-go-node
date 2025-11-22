import { calculateStatistics } from "../src/services/stats.service";

describe("calculateStatistics", () => {
  test("should compute max, min, sum, average and detect diagonal", () => {
    console.log("▶ Test: valid matrices");

    const matrices = [
      [
        [1, 0],
        [0, 3],
      ],
      [
        [5, 6],
        [7, 8],
      ],
    ];

    const stats = calculateStatistics(matrices);

    console.log("Computed stats:", stats);

    expect(stats.max).toBe(8);
    expect(stats.min).toBe(0);
    expect(stats.sum).toBe(1 + 0 + 0 + 3 + 5 + 6 + 7 + 8);
    expect(stats.average).toBeCloseTo(30 / 8);
    expect(stats.isDiagonal).toBe(true);
  });

  test("should detect non-diagonal matrices", () => {
    console.log("▶ Test: non-diagonal matrix");

    const matrices = [
      [
        [1, 2],
        [0, 3],
      ],
    ];

    const stats = calculateStatistics(matrices);

    console.log("Computed stats:", stats);

    expect(stats.isDiagonal).toBe(false);
  });

  test("should throw error when matrices are empty", () => {
    console.log("▶ Test: empty matrices");

    expect(() => calculateStatistics([])).toThrow("Empty matrices data");
  });

  test("should throw when matrices contain empty rows", () => {
    console.log("▶ Test: matrices with empty rows");

    const matrices = [
      [[]],
    ];

    expect(() => calculateStatistics(matrices)).toThrow();
  });
});
