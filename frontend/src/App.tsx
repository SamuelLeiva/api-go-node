import { useState } from "react";
import Login from "./pages/Login";
import QRPage from "./pages/QRPage";
import { getToken } from "./api/auth";

export default function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(!!getToken());

  if (!isLoggedIn) {
    return <Login onLoginSuccess={() => setIsLoggedIn(true)} />;
  }

  return <QRPage />;
}