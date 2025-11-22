import { Router } from "express";
import { processStats } from "../controllers/stats.controller";

const router = Router();

router.post("/", processStats);
export default router;
