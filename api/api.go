package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"khademi-practice/api/router"
	"khademi-practice/config"
)

func InitServer(cfg *config.Config, db *sqlx.DB) {
	r := gin.New()
	r.Use(gin.Recovery())
	RegisterRoutes(r, cfg, db)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config, db *sqlx.DB) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	{
		users := v1.Group("/users")

		router.User(users, cfg, db)
	}
}
