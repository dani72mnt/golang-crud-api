package userhandler

import (
	"github.com/gin-gonic/gin"
	validator "github.com/rezakhademix/govalidator/v2"
	"khademi-practice/dto"
	uservalidator "khademi-practice/validator/user"
	"net/http"
)

func (h UserHandler) Create(c *gin.Context) {
	var req dto.UserCreateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	vd := validator.New()
	validationErrors := uservalidator.UserValidator{}.ValidateCreateReq(vd, req)

	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	err := h.userSvc.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
