package main

import (
	"fmt"
	"net/http"

	"golangpkg/config"
)

func main(){
	config.ConnectDB()

	fmt.Println("Server is running...")
	http.ListenAndServe(":6060", nil)
}