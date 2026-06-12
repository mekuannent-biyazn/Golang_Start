package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json: "age"`
}

var students []Student 

func getUsers(w http.ResponseWriter, r *http.Request){
	if r.Method!= http.MethodGet{
		http.Error(
			w,
			"Method Not Allowed!",
			http.StatusBadRequest,
		)
	}
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	json.NewEncoder(w).Encode(&students)
}

func createStudent(w http.ResponseWriter, r *http.Request){
	var student Student

	if r.Method!= http.MethodPost{
		http.Error(
			w,
			"Method Not Allowed",
			http.StatusBadRequest,
		)
		return
	}

	err:= json.NewDecoder(
		r.Body,
	).Decode(&student)

	if student.Name==""{
		http.Error(
			w,
			"Student Name is requered!",
			http.StatusBadRequest,
		)
		return
	}

	if err!= nil {
		http.Error(
			w,
			"Invalid json",
			http.StatusBadRequest,
		)
		return
	}

	students = append(
		students, 
		student,
	)

	json.NewEncoder(w).Encode(&student)
	fmt.Fprintln(w, "New created student is!\n",student)
}

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello Well Come to Student Management API!!")
}

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case http.MethodGet:
			getUsers(w, r)
		case http.MethodPost:
			createStudent(w, r)
		default :
		http.Error(
			w,
			"Method Not Allowed!!",
			http.StatusBadRequest,
		)
		}
	},
)
	// http.HandleFunc("/user/register", createStudent)
	fmt.Println(`Server is running on http://localhost:7080`)
	http.ListenAndServe(":7080",nil)
}