package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
)

type AiController struct {
	Service services.AiService
}

func (controller *AiController) Prompt(c *gin.Context) {
	var payload types.PromptPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	reply := controller.Service.Prompt(payload.Prompt)
	c.JSON(http.StatusOK, reply)
}
