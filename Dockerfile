FROM goland:1.24.6-alpine as builder

WORKDIR /app

# Copia os arquivos de depend�ncias
COPY go.mod go.sum ./
# Baixa as depend�ncias
RUN go mod download

# Copia o resto do c�digo fonte
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