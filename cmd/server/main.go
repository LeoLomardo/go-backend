// cmd/server/main.go
package main

import (
	"go-backend/internal/database"
	"go-backend/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Conecta ao banco de dados
	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Cria um usuário admin padrão se ele não existir
	seedAdminUser()

	// Inicia a aplicação Fiber
	app := fiber.New()

	// Configura as rotas
	router.SetupRoutes(app)

	// Inicia o servidor na porta 8080
	log.Fatal(app.Listen(":8080"))
}

// Função para criar um usuário admin inicial para testes
func seedAdminUser() {
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username='admin')").Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if admin user exists: %v", err)
	}

	if !exists {
		// Usando o hash de senha (diferencial) [cite: 223]
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", "admin", string(hashedPassword))
		if err != nil {
			log.Fatalf("Failed to insert admin user: %v", err)
		}
		log.Println("Admin user created successfully")
	}
}
