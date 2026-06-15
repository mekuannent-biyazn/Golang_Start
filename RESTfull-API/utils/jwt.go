package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("super-secret-key")

func GenerateToken(userID int)(string, error){
	token:= jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user-ID": userID,
			"exp": time.Now().Add(5*time.Hour).Unix(),
		},
	)

	return token.SignedString(
		SecretKey,
	)
}