package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
	"github.com/lai0xn/hackiwna-backend/pkg/utils"
)

type AuthController struct {
	Service services.AuthService
}

func (controller *AuthController) Login(c *gin.Context) {
	var payload types.LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(payload.Email)
	user, err := controller.Service.Login(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := utils.GenerateToken(*user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (controller *AuthController) Signup(c *gin.Context) {
	var payload types.SignUpPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controller.Service.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, payload)
}
