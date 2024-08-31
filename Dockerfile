# Etapa de construção
FROM golang:1.20 AS builder

WORKDIR /app

# Copiar os arquivos de configuração e código
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend .

RUN go build -o main ./main.go

# Etapa de execução
FROM debian:bullseye-slim

WORKDIR /app

# Copiar o binário e os arquivos de configuração
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

# Definir o modo Gin para release
ENV GIN_MODE=release

EXPOSE 8080

CMD ["./main"]
