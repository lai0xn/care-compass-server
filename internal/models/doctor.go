package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DoctorProfile struct {
	gorm.Model
	Doctor    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DoctorID  uuid.UUID
	Field     string
	available bool
}
