package services

import (
	"github.com/lai0xn/hackiwna-backend/internal/models"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
	uuid "github.com/satori/go.uuid"
)

type UserService struct{}

func (s *UserService) ChangePfP() {
}

func (s *UserService) Update() {
}

func (s *UserService) UserByID(id uuid.UUID) (models.PatientProfile, error) {
	var user models.PatientProfile
	db := storage.DB.Preload("Patient").
		Where("patient_id = ?", id).
		First(&user)
	if db.Error != nil {
		return user, db.Error
	}
	return user, nil
}
