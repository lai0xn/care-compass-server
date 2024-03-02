package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/handlers"
	middlewares "github.com/lai0xn/hackiwna-backend/internal/transport/middlwares"
)

func aiRoutes(r *gin.Engine) {
	c := handlers.AiController{Service: services.AiService{}}
	aiR := r.Group("ai/chat")
	aiR.Use(middlewares.AuthMiddleware())
	aiR.POST("/", c.Prompt)
}
