import axios from "axios";
import { getToken } from "./auth";
import { API_GO_URL } from "../consts/urls";


export async function sendMatrix(matrix: number[][]) {
  const token = getToken();

  const res = await axios.post(
    `${API_GO_URL}/qr`,
    { matrix },
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  return res.data;
}
