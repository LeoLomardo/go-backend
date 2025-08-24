package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Conectado ao banco de dados com sucesso")
	createTables()
	return nil
}

func createTables() {

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, 
		username VARCHAR(50) UNIQUE NOT NULL, 
		password VARCHAR(255) NOT NULL
	);`

	_, err := DB.Exec(usersTable)
	if err != nil {
		panic("ERROR: Erro ao criar tabela user: " + err.Error())
	}

	campaignsTable := `
	CREATE TABLE IF NOT EXISTS campaigns (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		budget NUMERIC (12,2) NOT NULL,
		status VARCHAR(20) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`

	_, err = DB.Exec(campaignsTable)
	if err != nil {
		panic("ERROR: Eerro ao criar tabela campaigns: " + err.Error())
	}

	fmt.Println("Tabelas criadas com sucesso")
}
