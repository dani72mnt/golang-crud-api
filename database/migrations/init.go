package migrations

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func Up(db *sqlx.DB) {
	CreateUsersTable(db)
}

func CreateUsersTable(db *sqlx.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		family VARCHAR(100) NOT NULL,
		email VARCHAR(150) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	fmt.Println("Users table created successfully.")
}
