package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/handlers"
	middlewares "github.com/lai0xn/hackiwna-backend/internal/transport/middlwares"
)

func userRoutes(r *gin.Engine) {
	controller := handlers.UsersController{
		Service: services.UserService{},
	}
	a := r.Group("/users")
	a.Use(middlewares.AuthMiddleware())
	a.GET("/me", controller.Me)
	a.GET("/:id", middlewares.IsDoctor(), controller.GetUserByID)
}
