package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email       string    `gorm:"unique"`
	Password    string    `json:"-"`
	ProfilePic  string
	Name        string
	LastName    string
	PhoneNumber string
	Sex         bool
	IsActive    bool
	IsDoctor    bool
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV4()
	u.ProfilePic = "default.jpg"
	return nil
}
