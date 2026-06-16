package main

import (
	"net/http"
	"fmt"

	"go-auth-api/handlers"
	"go-auth-api/middleware"
)

func main(){
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/", handlers.GetAllUsers)
	http.HandleFunc("/profile", middleware.AuthMiddleware(handlers.Profile))
	http.HandleFunc("/updateprofile", middleware.AuthMiddleware(handlers.UpdateProfile))
	http.HandleFunc("/deleteuser", middleware.AuthMiddleware(handlers.DeleteUser))

	fmt.Println("server is running on localhost:5050")
	http.ListenAndServe(":5050", nil)
}