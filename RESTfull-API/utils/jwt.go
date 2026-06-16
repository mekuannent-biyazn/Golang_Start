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
			"user_id": userID,
			"exp": time.Now().Add(24*time.Hour).Unix(),
		},
	)

	return token.SignedString(
		SecretKey,
	)
}

func ValidateToken(
	tokenString string,
) (*jwt.Token, error) {

	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			return SecretKey, nil
		},
	)
}