# Étape 1 : Builder le projet Nuxt.js
FROM node:latest AS builder

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Étape 2 : Servir avec une image légère
FROM node:latest
WORKDIR /app
COPY --from=builder /app/.output ./
EXPOSE 3000
CMD ["node", "server/index.mjs"]
