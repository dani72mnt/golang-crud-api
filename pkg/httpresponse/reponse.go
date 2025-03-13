package httpresponse

import (
	"github.com/gin-gonic/gin"
)

type httpresponse struct {
	data any
}

func Send(c *gin.Context, data httpresponse, status int, err error) {
	c.JSON(status, data)
}
