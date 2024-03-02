package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/handlers"
)

func authRoutes(r *gin.Engine) {
	controller := handlers.AuthController{
		Service: *services.NewAuthService(),
	}
	auth := r.Group("/auth")
	auth.POST("/login", controller.Login)
	auth.POST("/signup", controller.Signup)
}
