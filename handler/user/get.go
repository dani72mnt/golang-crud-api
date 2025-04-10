package userhandler

import (
	"khademi-practice/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h UserHandler) GetAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid per_page"})
		return
	}

	users, err := h.userSvc.GetAll(c, page, perPage)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "there are no users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h UserHandler) Get(c *gin.Context) {
	var uri dto.UserIDURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

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
