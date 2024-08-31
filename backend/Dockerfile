# Etapa de construção
FROM golang:1.20-buster AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copie o arquivo go.mod e go.sum
COPY backend/go.mod backend/go.sum ./

# Baixe as dependências
RUN go mod download

# Copie o restante do código
COPY backend/ .

# Compile o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Etapa final
FROM debian:bullseye-slim

# Copie o binário da etapa de construção
COPY --from=builder /app/main /app/main

# Define o diretório de trabalho
WORKDIR /app

# Comando para iniciar o aplicativo
CMD ["./main"]
