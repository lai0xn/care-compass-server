package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/handlers"
)

func recordRoutes(r *gin.Engine) {
	recordService := &services.RecordService{}

	// Initialize controllers
	recordController := &handlers.RecordController{
		Service: recordService,
	}

	// Define routes

	records := r.Group("/records")
	{
		records.POST("/", recordController.CreateRecord)
		records.GET("/:id", recordController.GetRecordByID)
		records.PUT("/:id", recordController.UpdateRecord)
		records.DELETE("/:id", recordController.DeleteRecord)
	}
}
