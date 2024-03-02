package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PatientProfile struct {
	gorm.Model
	Patient          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PatientID        uuid.UUID
	NatID            string
	BloodGroup       string
	ChronicDisease   string
	ActiveMedication string
}
