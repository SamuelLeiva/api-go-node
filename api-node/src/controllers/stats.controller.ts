import { Request, Response } from "express";
import { calculateStatistics } from "../services/stats.service";

export function processStats(req: Request, res: Response) {
  const { matrices } = req.body;

  if (!matrices || !Array.isArray(matrices)) {
    console.log("❌ Invalid payload:", req.body);
    return res.status(400).json({ error: "matrices must be an array" });
  }

  try {
    const stats = calculateStatistics(matrices);
    console.log("📊 Computed statistics:", stats);
    return res.json(stats);
  } catch (err: any) {
    console.error("🔥 Error in /stats:", err);
    return res.status(500).json({ error: err.message });
  }
}
