import axios from "axios";
import { API_GO_URL } from "../consts/urls";

export async function login(username: string, password: string) {
  const res = await axios.post(`${API_GO_URL}/login`, {
    username,
    password,
  });

  const { token } = res.data;

  // Guardar el token
  localStorage.setItem("token", token);

  return token;
}

export function getToken() {
  return localStorage.getItem("token");
}
