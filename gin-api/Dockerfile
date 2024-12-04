# Étape 1 : Construire le binaire
FROM golang:latest AS builder

WORKDIR /app

# Copier les fichiers de configuration
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Compiler le binaire
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Étape 2 : Utiliser Debian Bookworm comme image finale
FROM debian:bookworm-slim

WORKDIR /app

# Installer les dépendances nécessaires
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copier le binaire
COPY --from=builder /app/main .

# Exposer le port utilisé par l'application
EXPOSE 8081

# Lancer le binaire
CMD ["./main"]
