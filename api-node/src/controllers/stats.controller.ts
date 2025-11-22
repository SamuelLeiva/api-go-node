import { Request, Response } from "express";
import { calculateStatistics } from "../services/stats.service";

export function processStats(req: Request, res: Response) {
  const { matrices } = req.body;

  if (!matrices || !Array.isArray(matrices)) {
    return res.status(400).json({ error: "matrices must be an array" });
  }

  try {
    const stats = calculateStatistics(matrices);
    return res.json(stats);
  } catch (err: any) {
    return res.status(500).json({ error: err.message });
  }
}
