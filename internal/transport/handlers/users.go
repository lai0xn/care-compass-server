package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	uuid "github.com/satori/go.uuid"
)

type UsersController struct {
	Service services.UserService
}

func (controller *UsersController) GetUserByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "not found")
		return
	}
	user, err := controller.Service.UserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "something wrong happened",
		})
		return
	}
	if user.Patient.Email == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, "not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UsersController) Me(c *gin.Context) {
	id, err := uuid.FromString(c.MustGet("id").(string))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	user, err := controller.Service.UserByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, user)
}
