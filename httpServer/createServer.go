package main

import (
	"fmt"
	"net/http"
	"io"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Well Come To Go Backend!!")
}

func About(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Well Come to About Page.")
}

func Contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Emaile: mekuannentbiyazn@gmail.com")
}

func Users(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Method)
}

func Methodvalidator(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(
			w, "Method Not Allowed!", http.StatusMethodNotAllowed,
		)
		return
	}
	fmt.Fprintln(w, "All user recoreds.")
}

func createUser(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	body, err:= io.ReadAll(r.Body)

	if err!= nil{
		http.Error(w, "Failed!", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "User created successfully!!\n", string(body))
}

func GerUserbyId(w http.ResponseWriter, r *http.Request){
	id:= r.URL.Query().Get("id")

	fmt.Fprintf(w, "User Id: %s", id)
}

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/users", GerUserbyId)
	http.HandleFunc("/methodvalidator",Methodvalidator)
	fmt.Println("Server is running on Port 8080")
	http.HandleFunc("/methodc",Users)
	http.ListenAndServe(":8080", nil)
}