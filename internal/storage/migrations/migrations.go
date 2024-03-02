package migrations

import (
	"fmt"

	"github.com/lai0xn/hackiwna-backend/internal/models"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
)

func Migrate() {
	fmt.Println("migrations")
	err := storage.DB.AutoMigrate(models.User{})
	if err != nil {
		panic(err)
	}
	storage.DB.AutoMigrate(models.Appointment{})
	storage.DB.AutoMigrate(models.DoctorProfile{})
	storage.DB.AutoMigrate(models.PatientProfile{})
}
