package main

import (
	"net/http"
	"fmt"

	"go-auth-api/handlers"
)

func main(){
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	http.Fprintln("server is running on localhost:5050")
	http.ListenAndServe(":5050", nil)
}