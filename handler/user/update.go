package userhandler

import (
	"github.com/gin-gonic/gin"
	validator "github.com/rezakhademix/govalidator/v2"
	"khademi-practice/dto"
	uservalidator "khademi-practice/validator/user"
	"net/http"
	"strconv"
)

func (h UserHandler) Update(c *gin.Context) {
	var req dto.UserUpdateReq
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	vd := validator.New()
	validationErrors := uservalidator.UserValidator{}.ValidateUpdateReq(vd, req)

	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	user, err := h.userSvc.Update(c, req, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
