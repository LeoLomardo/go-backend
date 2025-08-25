# Go Backend – Campaign Management API

Este projeto é uma **API REST** desenvolvida em **Go** com o framework **Fiber**, utilizando **PostgreSQL** como banco de dados e **JWT (JSON Web Tokens)** para autenticação.  
O sistema implementa um **CRUD de campanhas publicitárias** e gerenciamento de usuários, servindo como base sólida para aplicações de marketing, gestão de campanhas ou sistemas administrativos.

---

## Arquitetura do Projeto

A aplicação segue uma **arquitetura em camadas** com responsabilidades bem definidas:

```
cmd/
 ├── server/
 │    ├── main.go 
internal/
 ├── campaign/        → Módulo principal de campanhas
 │    ├── campaign.go          # Modelo (entidade Campaign)
 │    ├── repository.go        # Acesso ao banco de dados (DAO)
 │    ├── service.go           # Regras de negócio
 │    └── campaign_handler.go  # Controlador / HTTP Handlers
 |    └── service_test.go      # Testes da camada de Serviço
 │
 ├── user/            → Módulo de usuários
 │    ├── user.go             # Modelo de usuário
 │    └── handle_user.go      # Login e geração de JWT
 │
 ├── middleware/      → Middlewares
 │    └── auth.go            # Validação do token JWT
 │
 ├── database/        → Conexão e migração de tabelas
 │    └── db.go
 │
 └── router/          → Definição das rotas
      └── router.go
```

Essa separação em **Repository → Service → Handler** garante:
- **Repository**: apenas acesso ao banco, queries SQL e persistência.  
- **Service**: lógica de negócio, validações e orquestração.  
- **Handler**: lida com requisições HTTP, parse de payloads e respostas.  

---

## Principais Tecnologias e Bibliotecas

- **[Go](https://go.dev/)** – Linguagem principal
- **[Fiber](https://gofiber.io/)** – Framework web rápido e minimalista
- **[PostgreSQL](https://www.postgresql.org/)** – Banco de dados relacional
- **[JWT](https://jwt.io/)** – Autenticação baseada em tokens
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** – Hash seguro de senhas

---

## Funcionalidades Implementadas

### Autenticação
- Login via **username + senha**.  
- Senhas armazenadas com **bcrypt** (segurança reforçada).  
- Geração de **JWT válido por 72h**.  
- Middleware que protege rotas privadas, exigindo `Bearer Token`.  

### Campanhas (CRUD)
- **Criar campanha** – `POST /campaigns/`  
- **Listar todas campanhas** – `GET /campaigns/`  
- **Buscar por ID** – `GET /campaigns/:id`  
- **Atualizar** – `PUT /campaigns/:id`  
- **Deletar** – `DELETE /campaigns/:id`  

Cada campanha contém:
```json
{
  "id": 1,
  "name": "Campanha Exemplo",
  "budget": 5000.00,
  "status": "ativa",
  "created_at": "2025-08-25T12:00:00Z",
  "updated_at": "2025-08-25T12:00:00Z"
}
```

---

## Configuração do Banco de Dados

O projeto utiliza **variáveis de ambiente** para configuração:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=senha
DB_NAME=go_backend
```

Ao iniciar, o sistema cria automaticamente as tabelas `users` e `campaigns` caso não existam.

---

## Como Rodar a Aplicação

1. **Clonar o repositório**
   ```bash
   git clone https://github.com/seu-usuario/go-backend.git
   cd go-backend 
   ```

2. **Configurar variáveis de ambiente**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=senha
   export DB_NAME=go_backend
   ```

3. **Rodar a aplicação**
   ```bash
   sudo systemctl start docker
   sudo docker compose up --build 
   ```

4. **Testar endpoints**
   Para testar os endpoints da aplicação, utilizei o software **Postman**

---

## Decisões de Design

- **Fiber** foi escolhido por exigencias do projeto.  
- **Arquitetura em camadas (Repository, Service, Handler)** facilita manutenção, testes unitários e extensibilidade.  
- **JWT + bcrypt** garantem um fluxo de autenticação seguro.  
- **PostgreSQL** foi selecionado por exigencias do projeto.

---



## Autor

Desenvolvido por **Leo Lomardo**  
