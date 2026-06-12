package main

import (
	"fmt"
	"net/http"
	// "io"
	"encoding/json"
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

type Usersj struct{
	Id int `json: "id"`
	Name string `json: "name"`
}

func createUser(w http.ResponseWriter, r *http.Request){
	var user Usersj

	erru:= json.NewDecoder(
		r.Body,
	).Decode(&user)

	if erru!= nil {
		http.Error(
			w,
			"Invalid request!",
			http.StatusBadRequest,
		)
		return
	}
	fmt.Fprintln(w,"Id: ", user.Id)
	fmt.Fprintln(w,"Name: ", user.Name)

	// w.WriteHeader(http.StatusCreated)
	// body, err:= io.ReadAll(r.Body)

	// if err!= nil{
	// 	http.Error(w, "Failed!", http.StatusBadRequest)
	// 	return
	// }

	// fmt.Fprintln(w, "User created successfully!!\n", string(body))
}

func GerUserbyId(w http.ResponseWriter, r *http.Request){
	id:= r.URL.Query().Get("id")

	fmt.Fprintf(w, "User Id: %s", id)
}

func getUsers(w http.ResponseWriter, r *http.Request){
	users:= []Usersj{
		{1, "mekuannent"},
		{2, "Tirusiet"},
	}

	jsonData,_:= json.MarshalIndent(
		users,
		"",
		" ",
	) 

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	user1:= Usersj{
		Id: 3,
		Name: "Hiemu",
	}
	json.NewEncoder(w).Encode(user1)
	fmt.Fprintln(w,"All Registerd Users Are!\n", string(jsonData))
}

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/users/all",getUsers)
	http.HandleFunc("/users", GerUserbyId)
	http.HandleFunc("/methodvalidator",Methodvalidator)
	fmt.Println("Server is running on Port 8080")
	http.HandleFunc("/methodc",Users)
	http.ListenAndServe(":8080", nil)
}