package services

import (
	"errors"

	"github.com/lai0xn/hackiwna-backend/internal/models"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
	"github.com/lai0xn/hackiwna-backend/pkg/utils"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(user types.SignUpPayload) error {
	userModel := &models.User{
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    utils.HashPassword(user.Password),
		PhoneNumber: user.PhoneNumber,
		IsActive:    user.IsActive,
		IsDoctor:    user.IsDoctor,
	}
	db := storage.DB.Create(userModel)
	if db.Error != nil {
		return db.Error
	}
	if user.IsDoctor {
		storage.DB.Create(&models.DoctorProfile{
			DoctorID: userModel.ID,
		})
	} else {
		storage.DB.Create(&models.PatientProfile{
			PatientID: userModel.ID,
		})
	}
	return nil
}

func (s *AuthService) Login(payload types.LoginPayload) (*models.User, error) {
	var user models.User
	storage.DB.Model(&user).Where("email = ?", payload.Email).First(&user)
	if user.Email == "" {
		return nil, errors.New("Invalid login")
	}
	if !utils.CheckPassword(payload.Password, user.Password) {
		return nil, errors.New("Invalid login")
	}
	return &user, nil
}
