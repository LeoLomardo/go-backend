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

	fmt.Println("Database connection successful")
	createTables()
	return nil
}

func createTables() {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, [cite: 128]
		username VARCHAR(50) UNIQUE NOT NULL, [cite: 129]
		password VARCHAR(255) NOT NULL [cite: 130]
	);`

	_, err := DB.Exec(usersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	campaignsTable := `
	CREATE TABLE IF NOT EXISTS campaigns (
		id SERIAL PRIMARY KEY, [cite: 133]
		name VARCHAR(255) NOT NULL, [cite: 139]
		budget NUMERIC (12,2) NOT NULL, [cite: 140]
		status VARCHAR(20) NOT NULL, [cite: 140]
		created_at TIMESTAMP DEFAULT NOW(), [cite: 140]
		updated_at TIMESTAMP DEFAULT NOW() [cite: 140]
	);`
	_, err = DB.Exec(campaignsTable)
	if err != nil {
		panic("Could not create campaigns table: " + err.Error())
	}

	fmt.Println("Tables created successfully")
}
