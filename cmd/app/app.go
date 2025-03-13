package app

import (
	"khademi-practice/config"
	userhandler "khademi-practice/handler/user"
	userrepository "khademi-practice/repository"
	userservice "khademi-practice/service/user"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	Cfg *config.Config
	DB  *sqlx.DB

	// make it better
	UserService userservice.UserService
	UserHandler userhandler.UserHandler
}

func New(cfg *config.Config, db *sqlx.DB) *Application {
	userSvc := userservice.New(cfg, userrepository.New(db))

	return &Application{
		UserService: userSvc,
		UserHandler: userhandler.New(userSvc),
	}
}
