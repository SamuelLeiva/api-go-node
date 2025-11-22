import axios from "axios";

describe("FULL INTEGRATION TEST: Go → Node → stats", () => {
  const goUrl = "http://localhost:8080/qr"; // expuesto por docker compose
  
  test("Should compute QR and propagate results to Node API", async () => {
    console.log("▶ Sending matrix to API-Go...");

    const payload = {
      matrix: [
        [1, 2],
        [3, 4],
      ],
    };

    const response = await axios.post(goUrl, payload);
    
    console.log("QR Response from Go:", response.data);

    const { q, r } = response.data;

    expect(q).toBeDefined();
    expect(r).toBeDefined();

    console.log("Checking Node logs for forwarded stats...");

    // Now query the Node API directly to verify stats
    const nodeResponse = await axios.post(
      "http://localhost:3000/stats",
      { matrices: [q, r] }
    );

    console.log("Stats response from Node:", nodeResponse.data);

    expect(nodeResponse.data.max).toBeDefined();
    expect(nodeResponse.data.sum).toBeGreaterThan(0);
  });
});
