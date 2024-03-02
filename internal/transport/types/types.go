package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SignUpPayload struct {
	Name        string `json:"name"        binding:"required"`
	LastName    string `json:"lastName"    binding:"required"`
	Password    string `json:"password"    binding:"required"`
	Email       string `json:"email"       binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	IsDoctor    bool   `json:"isDoctor"`
	IsActive    bool   `json:"isActive"`
	Sex         bool   `json:"sex"`
}

type LoginPayload struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AppointmentPayload struct {
	Date      time.Time `json:"date"      binding:"required"`
	PatientID uuid.UUID `json:"patientId" binding:"required"`
	DoctorID  uuid.UUID `json:"doctorId"  binding:"required"`
}

type PromptPayload struct {
	Prompt string `json:"prompt" binding:"required"`
}
type RecordPayload struct {
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	DoctorID  uuid.UUID `json:"doctor_id"`
	PatientID uuid.UUID `json:"patient_id"`
	Date      time.Time `json:"date"`
	File      string    `json:"file"`
}
