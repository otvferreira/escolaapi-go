# Use uma imagem base para construção
FROM golang:1.20-buster AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o código fonte e o go.mod
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Construa o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Use uma imagem base leve para o runtime
FROM debian:bullseye-slim

# Copie o binário do estágio de construção
COPY --from=builder /app/main /app/main

# Defina o diretório de trabalho
WORKDIR /app

# Defina o comando para rodar o binário
CMD ["./main"]

