package services

import (
	"time"

	"github.com/lai0xn/hackiwna-backend/internal/models"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
	uuid "github.com/satori/go.uuid"
)

type RecordService struct{}

// CreateRecord creates a new record with optional file upload
func (s *RecordService) CreateRecord(req types.RecordPyload) (*models.Record, error) {
	record := &models.Record{
		Title:     req.Title,
		Summary:   req.Summary,
		DoctorID:  req.DoctorID,
		PatientID: req.PatientID,
		Date:      req.Date,
		File:      req.File,
	}

	if err := storage.DB.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// GetRecordByID retrieves a record by ID
func (s *RecordService) GetRecordByID(id uuid.UUID) (*models.Record, error) {
	var record models.Record
	if err := storage.DB.Preload("Doctor").Preload("Patient").Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// UpdateRecord updates an existing record
func (s *RecordService) UpdateRecord(id uuid.UUID, req types.RecordPyload) error {
	var record models.Record
	if err := storage.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}

	record.Title = req.Title
	record.Summary = req.Summary
	record.Date = req.Date
	record.File = req.File

	if err := storage.DB.Save(&record).Error; err != nil {
		return err
	}

	return nil
}

// DeleteRecord deletes an existing record
func (s *RecordService) DeleteRecord(id uuid.UUID) error {
	var record models.Record
	if err := storage.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}

	if err := storage.DB.Delete(&record).Error; err != nil {
		return err
	}

	return nil
}
