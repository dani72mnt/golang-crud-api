package main

import (
	"khademi-practice/api"
	"khademi-practice/cmd/app"
	"khademi-practice/config"
	db "khademi-practice/database"
	"khademi-practice/database/migrations"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	cfg := config.LoadConfig()

	connectDb, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatalf("Error connecting to connectDb: %v", err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(connectDb)

	boil.SetDB(connectDb)
	migrations.Up(connectDb)
	log.Println("Migrations completed successfully.")

	app := app.New(cfg, connectDb)

	api.SetupServer(app)
}
