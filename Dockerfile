# Etapa de build
FROM golang:1.20 AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o go.mod e go.sum para o diretório de trabalho
COPY backend/go.mod backend/go.sum ./

# Baixe as dependências
RUN go mod download

# Copie o restante do código para o diretório de trabalho
COPY backend .

# Compile o binário
RUN go build -o main ./main.go

# Etapa final: usar uma imagem mínima para rodar a aplicação
FROM debian:bullseye-slim

# Copie o binário da etapa de build
COPY --from=builder /app/main /app/main

# Defina o diretório de trabalho
WORKDIR /app

# Exponha a porta em que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/app/main"]

