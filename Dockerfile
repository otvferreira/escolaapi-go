# Use uma imagem base do Go
FROM golang:1.20 AS builder

# Defina o diretório de trabalho
WORKDIR /backend

# Copie o código-fonte para o diretório de trabalho
COPY . .

# Compile a aplicação
RUN go build -o main .

# Use uma imagem base menor para rodar a aplicação
FROM debian:bullseye-slim

# Copie o binário da imagem de build
COPY --from=builder /backend/main /backend/main

# Defina o comando para rodar a aplicação
CMD ["/backend/main"]
