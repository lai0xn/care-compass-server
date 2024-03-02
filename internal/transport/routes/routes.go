package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "hello world")
	})

	// Auth routes
	authRoutes(r)

	// Appointment routes

	appointmentRoutes(r)

	// Ai routes
	aiRoutes(r)

	// users routes
	userRoutes(r)

	// records routes

	recordRoutes(r)
}
