package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid"`
	Summary   string
	Doctor    User
	Patient   User
	Date      time.Time
	DoctorID  uuid.UUID
	PatientID uuid.UUID
}
