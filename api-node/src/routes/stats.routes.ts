import { Router } from "express";
import { processStats } from "../controllers/stats.controller";
import { authRequired } from "../middleware/auth.middleware";

const router = Router();

router.post("/", authRequired, processStats);

export default router;
