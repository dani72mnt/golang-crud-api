package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"khademi-practice/api"
	"khademi-practice/cmd/app"
	"khademi-practice/config"
	db "khademi-practice/database"
	"log"
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

	application := app.New(cfg, connectDb)

	api.SetupServer(application)
}
