# 📘 QR Matrix Processing – Proyecto Fullstack

Este proyecto implementa un sistema completo que:

1. Recibe una **matriz** en la API en Go  
2. Calcula su **descomposición QR (Q y R)**  
3. Envía automáticamente los resultados a la API en Node.js  
4. La API en Node calcula estadísticas globales (max, min, suma, promedio, diagonalidad)  
5. Un frontend en React muestra todo en la pantalla con autenticación básica (JWT)

---

## 🚀 Tecnologías

- **Go + Fiber** → API principal para descomposición QR  
- **Node.js + Express + TypeScript** → API secundaria para estadísticas  
- **React + Vite + Tailwind** → Frontend  
- **Docker + Docker Compose** → Orquestación  
- **JWT** → Autenticación  
- **CORS** → Control de dominios permitidos  

---

## Flujo

1. Frontend → `/login`  
2. Node genera JWT  
3. Frontend lo guarda en `localStorage`  
4. Todas las peticiones llevan el header Authorization: Bearer token

---

## 🐳 Cómo levantar el proyecto con Docker

Desde la raíz:

```bash
    docker-compose up --build
```

---

## 🧭 Rutas del Proyecto

## Frontend (React)

```bash
 http://localhost:5173/
 http://44.200.73.216:5173/
```

## API-Go

### Login (POST)

```bash
 http://localhost:8080/login
 http://44.200.73.216:8080/login
```

Request

```json
    {
  "username": "admin",
  "password": "123456"
}
```

Response

```json
    {
  "token": "<jwt>"
}
```

### ProcessQR (POST)

```bash
 http://localhost:8080/qr
 http://44.200.73.216:8080/qr
```

Request

```json
    {
    "matrix": [
        [
            1,
            2
        ],
        [
            3,
            4
        ]
    ]
}
```

Response

```json
    {
    "Q": [
        [
            -0.316227766016838,
            -0.9486832980505138
        ],
        [
            -0.9486832980505138,
            0.316227766016838
        ]
    ],
    "R": [
        [
            -3.1622776601683795,
            -4.427188724235731
        ],
        [
            0,
            -0.6324555320336751
        ]
    ],
    "stats": {
        "average": -1.2649110640673515,
        "isDiagonal": false,
        "max": 0.316227766016838,
        "min": -4.427188724235731,
        "sum": -10.119288512538812
    }
}
```

## API-Node

### Process Stats (POST)

```bash
 http://localhost:3000/stats
 http://44.200.73.216:3000/stats
```

Request

```json
    {
    "matrices": [
        [
            [
                1,
                0
            ],
            [
                0,
                1
            ]
        ],
        [
            [
                5,
                6
            ],
            [
                7,
                8
            ]
        ]
    ]
}
```

Response

```json
    {
    "max": 8,
    "min": 0,
    "sum": 28,
    "average": 3.5,
    "isDiagonal": true
}
```

---

## Notas finales

- El frontend siempre envía el JWT automáticamente

- La API Go reenvía Q y R a Node
