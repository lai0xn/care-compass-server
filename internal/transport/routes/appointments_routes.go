package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/handlers"
	middlewares "github.com/lai0xn/hackiwna-backend/internal/transport/middlwares"
)

func appointmentRoutes(r *gin.Engine) {
	controller := handlers.AppointmentController{Service: services.AppointmentService{}}
	a := r.Group("/appointments")
	a.Use(middlewares.AuthMiddleware())
	a.GET("/me", controller.Get)
	a.POST("/set", controller.Post)
}
