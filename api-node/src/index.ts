import express from "express";
import statsRoutes from "./routes/stats.routes";

const app = express();
app.use(express.json());

app.use("/stats", statsRoutes);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`API-Node running on port ${PORT}`);
});
