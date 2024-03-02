package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	}
	return string(hashed_password)
}

func CheckPassword(password string, hashed_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
