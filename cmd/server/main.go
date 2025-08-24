package main

import (
	"go-backend/internal/database"
	"go-backend/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatalf("ERROR: Falha ao conectar com o banco de dados: %v", err)
	}

	seedAdminUser()

	app := fiber.New()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

// Funcao q cria um usuario admin inicial caso nao exista nenhum outro
func seedAdminUser() {
	var exists bool

	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username='admin')").Scan(&exists)
	if err != nil {
		log.Fatalf("ERROR: Falha na checagem se o usuario existe: %v", err)
	}

	if !exists {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("ERROR: Falha ao criar hash da senha: %v", err)
		}

		_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", "admin", string(hashedPassword))
		if err != nil {
			log.Fatalf("ERROR: Falha em inserir usuario admin: %v", err)
		}

		log.Println("Admin criado com sucesso")
	}
}
