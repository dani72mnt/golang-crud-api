package db

import (
	"fmt"
	"khademi-practice/config"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDb(cfg *config.Config) (*sqlx.DB, error) {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User,
		cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode,
	)

	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Cannot connect to Postgres DB: %v", err)
		return nil, err
	}

	DB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	DB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	DB.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)

	err = DB.Ping()
	if err != nil {
		log.Fatal(fmt.Errorf("database is unreachable, err: %w", err)) // wrap
		return nil, err
	}

	fmt.Println("Connected to Postgres DB")

	return DB, nil
}
