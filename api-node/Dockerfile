FROM node:20-alpine

WORKDIR /app

COPY package.json package-lock.json ./

# Evitar errores de permisos y lockfile rotos
RUN npm config set fund false \
    && npm config set update-notifier false

RUN npm install

COPY . .

# Asegurar permisos correctos
RUN chmod -R +x node_modules/.bin

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
