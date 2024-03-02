package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/hackiwna-backend/internal/models"
)

var jwtSecret = "this is the secrey key"

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
		"Id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	return token
}

func RefreshToken(refreshToken string) (string, error) {
	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	// Check if token is valid and not expired
	if !token.Valid {
		return "", errors.New("Invalid token")
	}

	// Extract user information from the token, such as user ID or username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid token claims")
	}

	// Generate a new JWT token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims["user_id"],                       // Assuming user ID is stored in the refresh token claims
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // Token expiration time
	})

	// Sign the token with the secret key
	tokenString, err := newToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
