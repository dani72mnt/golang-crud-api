package router

import (
	"khademi-practice/cmd/app"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, app *app.Application) {
	router.GET("/", app.UserHandler.GetAll)
	//router.GET("/:id", app.UserHandler.Get)
	router.POST("/", app.UserHandler.Create)
	//router.PATCH("/:id", app.UserHandler.Update)
	//router.DELETE("/:id", app.UserHandler.Delete)
}
