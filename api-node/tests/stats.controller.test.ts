import { processStats } from "../src/controllers/stats.controller";
import * as statsService from "../src/services/stats.service";

// mock del servicio
jest.mock("../src/services/stats.service");

describe("processStats controller", () => {
  const mockRequest = (body: any) => ({
    body,
  }) as any;

  const mockResponse = () => {
    const res: any = {};
    res.status = jest.fn().mockReturnValue(res);
    res.json = jest.fn().mockReturnValue(res);
    return res;
  };

  test("should return 400 if matrices is missing", () => {
    console.log("▶ Controller Test: invalid payload");

    const req = mockRequest({});
    const res = mockResponse();

    processStats(req, res);

    expect(res.status).toHaveBeenCalledWith(400);
    expect(res.json).toHaveBeenCalledWith({
      error: "matrices must be an array",
    });
  });

  test("should return 200 with stats result", () => {
    console.log("▶ Controller Test: valid payload");

    const req = mockRequest({
      matrices: [
        [
          [1, 0],
          [0, 2],
        ],
      ],
    });

    const res = mockResponse();

    (statsService.calculateStatistics as jest.Mock).mockReturnValue({
      max: 2,
      min: 0,
      sum: 3,
      average: 1.5,
      isDiagonal: true,
    });

    processStats(req, res);

    expect(res.status).not.toHaveBeenCalled();
    expect(res.json).toHaveBeenCalledWith({
      max: 2,
      min: 0,
      sum: 3,
      average: 1.5,
      isDiagonal: true,
    });
  });

  test("should return 500 when service throws an error", () => {
    console.log("▶ Controller Test: service throws error");

    const req = mockRequest({
      matrices: [[[1]]],
    });
    const res = mockResponse();

    (statsService.calculateStatistics as jest.Mock).mockImplementation(() => {
      throw new Error("Service failure");
    });

    processStats(req, res);

    expect(res.status).toHaveBeenCalledWith(500);
    expect(res.json).toHaveBeenCalledWith({
      error: "Service failure",
    });
  });
});
