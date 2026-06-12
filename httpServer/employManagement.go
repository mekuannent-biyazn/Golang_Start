package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Employee struct{
	ID int `json: "id"`
	Name string `json: "name"`
	Salary float64 `json: salary`
}

var employees []Employee

func createEmployee(w http.ResponseWriter, r *http.Request){
	var employee Employee

	if r.Method != http.MethodPost{
		http.Error(
			w,
			"Method Not Allowed!!",
			http.StatusBadRequest,
		)
		return
	}
	
	err:= json.NewDecoder(
		r.Body,
	).Decode(&employee)

	if employee.Name==""{
		http.Error(
			w,
			"Employee Name is Requered!!",
			http.StatusBadRequest,
		)
		return
	}

	if employee.Salary <=5000{
		http.Error(
			w,
			"Employee Salary Must Be Greater than 5000!",
			http.StatusBadRequest,
		)
		return
	}

	if err!= nil{
		http.Error(
			w,
			"Invalid Information",
			http.StatusBadRequest,
		)
	}

	employees= append(employees, employee)

	json.NewEncoder(w).Encode(&employee)
}

func getEmployees(w http.ResponseWriter, r *http.Request){
	if r.Method!= http.MethodGet{
		http.Error(
			w,
			"Method Not Alloewd!!",
			http.StatusBadRequest,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(&employees)
}


func main(){
	http.HandleFunc("/employees",func(
		w http.ResponseWriter,
		r *http.Request,
	){
		switch r.Method{
		case http.MethodPost:
			createEmployee(w, r)
		case http.MethodGet:
			getEmployees(w, r)
		default :
		http.Error(
			w,
			"Method Not Allowed!!",
			http.StatusBadRequest,
		)
		}
	},
)
	fmt.Println("Server is running on http://locaalhost:8090")
	http.ListenAndServe(":8090", nil)
}