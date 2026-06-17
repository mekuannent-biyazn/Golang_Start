package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey []byte

func LoadJWTSecret() {
	SecretKey = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user-ID": userID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		},
	)
	return token.SignedString(SecretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
}