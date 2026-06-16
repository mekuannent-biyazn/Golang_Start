package middleware

import (
	"net/http"
	"strings"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"fmt"

	"go-auth-api/utils"
)

func AuthMiddleware(
	next http.HandlerFunc,
)http.HandlerFunc{
	return func(
		w http.ResponseWriter,
		r *http.Request,
	){
		authHeader:= r.Header.Get(
			"Authorization",
		)

		if authHeader== "" {
			http.Error(
				w,
				"Token error!",
				http.StatusUnauthorized,
			)
			return
		}

		parts := strings.Split(
			authHeader,
			" ",
		)

		if len(parts) !=2{
			http.Error(
				w,
				"Token error!",
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := parts[1]

		token, err := utils.ValidateToken(
			tokenString,
		)

		if err != nil || !token.Valid{
			http.Error(
				w,
				"Token error!",
				http.StatusUnauthorized,
			)
			return
		}

		// jwt.MapClaims{
		// 	"user_id": userID,
		// }

		claims:= token.Claims.(jwt.MapClaims)

		fmt.Println("Claims:", claims)

		userID :=
		int(
			claims["user_id"].
				(float64),
		)

		fmt.Println("UserID:", userID)

		fmt.Println("Saving userID into Context:", userID)
		ctx := context.WithValue(
			r.Context(),
			"userID",
			userID,
		)
		r = r.WithContext(ctx)

		next(w, r)
	}
}