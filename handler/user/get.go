package userhandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h UserHandler) GetAll(c *gin.Context) {
	users, err := h.userSvc.GetAll(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "there are no users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h UserHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userSvc.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
