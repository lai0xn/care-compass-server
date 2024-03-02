package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Patient   User
	Doctor    User
	DoctorID  uuid.UUID
	PatientID uuid.UUID
	Date      time.Time
}

func (a *Appointment) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.NewV4()
	return nil
}
