import { useState } from "react";
import { sendMatrix } from "../api/matrix";

export default function QRPage() {
  const [matrixInput, setMatrixInput] = useState("[[1,2],[3,4]]");
  const [result, setResult] = useState<any>(null);

  async function handleSend() {
    try {
      const parsed = JSON.parse(matrixInput);
      const res = await sendMatrix(parsed);
      setResult(res);
    } catch (err) {
      alert("Error sending matrix");
    }
  }

  return (
    <div style={{ margin: "50px" }}>
      <h2>QR Processing</h2>

      <textarea
        value={matrixInput}
        onChange={(e) => setMatrixInput(e.target.value)}
        rows={4}
        style={{ width: "400px" }}
      />

      <br />
      <button onClick={handleSend} style={{ marginTop: "20px" }}>
        Enviar Matriz
      </button>

      {result && (
        <pre style={{ background: "#eee", padding: "20px", marginTop: "20px" }}>
          {JSON.stringify(result, null, 2)}
        </pre>
      )}
    </div>
  );
}
