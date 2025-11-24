import express from "express";
import statsRoutes from "./routes/stats.routes";
import dotenv from "dotenv";
import path from "path";
import cors from "cors";

dotenv.config({
  path: path.resolve(__dirname, "../.env"),
})

const app = express();

app.use(cors({
  origin: '*',
  methods: ['GET', 'POST', 'OPTIONS'],
  allowedHeaders: ['Origin', 'Content-Type', 'Accept', 'Authorization'],
}))

app.use(express.json());

app.use("/stats", statsRoutes);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`API-Node running on port ${PORT}`);
});
