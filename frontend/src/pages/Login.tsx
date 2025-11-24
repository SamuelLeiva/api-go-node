import { useState } from "react";
import { login } from "../api/auth";

export default function Login({ onLoginSuccess }: { onLoginSuccess: () => void }) {
  const [username, setUsername] = useState("admin");
  const [password, setPassword] = useState("123456");
  const [error, setError] = useState("");

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");

    try {
      await login(username, password);
      onLoginSuccess();
    } catch (err) {
      setError("Invalid credentials");
    }
  }

  return (
    <div style={{ margin: "50px auto", width: "300px" }}>
      <h2>Login</h2>

      {error && <p style={{ color: "red" }}>{error}</p>}

      <form onSubmit={handleSubmit}>
        <div>
          <label>User:</label>
          <input 
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            style={{ width: "100%" }}
          />
        </div>

        <div>
          <label>Password:</label>
          <input 
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ width: "100%" }}
          />
        </div>

        <button type="submit" style={{ marginTop: "20px", width: "100%" }}>
          Login
        </button>
      </form>
    </div>
  );
}
