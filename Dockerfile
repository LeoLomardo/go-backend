FROM golang:1.24.6-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila a aplica��o
# O -o server vai gerar um execut�vel chamado "server" dentro da pasta /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/server/server ./cmd/server

# Est�gio final (imagem final, menor)
FROM alpine:latest

WORKDIR /root/

# Copia o execut�vel compilado do est�gio de build
COPY --from=builder /app/cmd/server/server .

# Exp�e a porta que a aplica��o vai usar
EXPOSE 8080

# Comando para rodar a aplica��o quando o container iniciar
CMD ["./server"]