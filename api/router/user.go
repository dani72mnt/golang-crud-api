package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"khademi-practice/config"
	userhandler "khademi-practice/handler/user"
	userrepository "khademi-practice/repository"
	userservice "khademi-practice/service/user"
)

func User(router *gin.RouterGroup, cfg *config.Config, db *sqlx.DB) {

	repo := userrepository.New(db)
	srv := userservice.New(cfg, repo)
	h := userhandler.New(srv)

	router.GET("/", h.GetAll)
	router.GET("/:id", h.Get)
	router.POST("/", h.Create)
	router.PATCH("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
}
