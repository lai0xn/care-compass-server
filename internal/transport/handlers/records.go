package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/services"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
	uuid "github.com/satori/go.uuid"
)

type RecordController struct {
	Service *services.RecordService
}

// CreateRecord handles POST request to create a new record with optal file upload
func (ctrl *RecordController) CreateRecord(c *gin.Context) {
	// Parse the form file
	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if file != nil {
		// Save the file to a locat
		filePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
	}

	var payload types.RecordPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := ctrl.Service.CreateRecord(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func (ctrl *RecordController) GetRecordByID(c *gin.Context) {
	// Parse record ID from URL parameter
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	// Fetch record from service
	record, err := ctrl.Service.GetRecordByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	// Respond with record data
	c.JSON(http.StatusOK, record)
}

func (ctrl *RecordController) DeleteRecord(c *gin.Context) {
	// Parse record ID from URL parameter
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	// Delete record from service
	err = ctrl.Service.DeleteRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}

// UpdateRecord handles PUT request to update an existing record with optal file upload
func (ctrl *RecordController) UpdateRecord(c *gin.Context) {
	// Parse the form file
	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if file != nil {
		// Save the file to a locat
		filePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
	}

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	var payload types.RecordPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.Service.UpdateRecord(id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
