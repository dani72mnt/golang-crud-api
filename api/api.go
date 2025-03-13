package api

import (
	"fmt"
	"khademi-practice/api/router"
	"khademi-practice/cmd/app"

	"github.com/gin-gonic/gin"
)

func SetupServer(app *app.Application) {
	r := gin.New()
	r.Use(gin.Recovery())

	RegisterRoutes(r, app)

	err := r.Run(fmt.Sprintf(":%s", app.Cfg.Server.Port))
	if err != nil {
		return
	}
}

func RegisterRoutes(r *gin.Engine, app *app.Application) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	{
		users := v1.Group("/users")
		router.User(users, app)
	}
}
