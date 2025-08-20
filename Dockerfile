FROM goland:1.24.6-alpine as builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
# Baixa as dependências
RUN go mod download

# Copia o resto do código fonte
COPY . .

# Compila a aplicação
# O -o server vai gerar um executável chamado "server" dentro da pasta /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/server/server ./cmd/server

# Estágio final (imagem final, menor)
FROM alpine:latest

WORKDIR /root/

# Copia o executável compilado do estágio de build
COPY --from=builder /app/cmd/server/server .

# Expõe a porta que a aplicação vai usar
EXPOSE 8080

# Comando para rodar a aplicação quando o container iniciar
CMD ["./server"]