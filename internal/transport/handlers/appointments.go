package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
	uuid "github.com/satori/go.uuid"
)

type AppointmentController struct {
	Service services.AppointmentService
}

func (controller *AppointmentController) Get(c *gin.Context) {
	id := c.MustGet("id").(string)
	apps := controller.Service.GetAppointments(uuid.FromStringOrNil(id))
	c.JSON(http.StatusOK, gin.H{
		"appointments": apps,
	})
}

func (controller *AppointmentController) Post(c *gin.Context) {
	var payload types.AppointmentPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return

	}
	err = controller.Service.CreateApointment(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "appointment created",
	})
}
