package main

import (
	"fmt"
	"net/http"

	"golangpkg/config"
	"golangpkg/handlers"
	"golangpkg/middleware"
	"golangpkg/utils"
)

func main() {
	config.LoadEnv()
	utils.LoadJWTSecret()
	config.ConnectDB()

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/profile", middleware.AuthMiddleware(handlers.GetProfile))

	fmt.Println("Server is running on port 6060...")
	fmt.Println("Available routes:")
	fmt.Println("  POST /register - Register a new user")
	fmt.Println("  POST /login - Login user")
	fmt.Println("  GET /profile - Get user profile (requires authentication)")
	
	http.ListenAndServe(":6060", nil)
}