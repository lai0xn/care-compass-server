package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Record struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Title   string
	Summary string

	DoctorID  uuid.UUID `gorm:"type:uuid"`
	PatientID uuid.UUID `gorm:"type:uuid"`
	Date      time.Time

	File      string // File path or name
	CreatedAt time.Time
	UpdatedAt time.Time

	Doctor  User `gorm:"foreignkey:DoctorID"`
	Patient User `gorm:"foreignkey:PatientID"`
}

func (r *Record) BeforeCreate(tx gorm.DB) error {
	r.ID = uuid.NewV4()
	return nil
}
