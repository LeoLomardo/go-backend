FROM golang:1.24.6-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila a aplicao
# O -o server vai gerar um executavel chamado "server" dentro da pasta /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/server/server ./cmd/server

FROM alpine:latest

WORKDIR /root/

# Copia o executavel compilado do estagio de build
COPY --from=builder /app/cmd/server/server .

EXPOSE 8080

# Comando para rodar a aplicao quando o container iniciar
CMD ["./server"]