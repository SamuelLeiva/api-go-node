# ========================
# 1) Build Stage
# ========================
FROM node:20-alpine AS builder

WORKDIR /app

ARG VITE_API_GO_URL
ENV VITE_API_GO_URL=${VITE_API_GO_URL}

# Copiar dependencias
COPY package.json package-lock.json ./
RUN npm install

# Copiar el resto del código
COPY . .

# Construir producción estática
RUN npm run build


# ========================
# 2) Serve Stage (nginx)
# ========================
FROM nginx:alpine

# Copiar el build al directorio público de nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Reemplazar configuración nginx (SOLO si necesitas CORS o rutas SPA)
# COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]