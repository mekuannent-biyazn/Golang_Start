package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct{
	ID int `json:"id"`
	Name string `json: "name"`
	Phone string `json: "phone"`
}

type Driver struct{
	ID int `json: "id"`
	Name string `json: "name"`
	CarNo string `json: "carno"`
}

func main(){
	// struct to json converssion Marshaling
	driver:= Driver{
		ID: 10,
		Name: "Zelalem",
		CarNo: "12345: AM",
	}

	jsonCD, erro:= json.MarshalIndent(
		driver,
		"",
		" ",
	)

	if erro != nil {
		fmt.Println(erro)
		return
	}
	fmt.Println(string(jsonCD))
 
	fmt.Println("=================================================")




	crData:=[]byte(`
	{
	"ID": 11,
	"Name": "alemu",
	"CarNo": "AM: 43234"
	}
	`)


	var drrive Driver 

	errd:= json.Unmarshal(crData, &drrive)
	if errd!= nil{
		fmt.Println(errd)
		return
	}
	fmt.Println("Id: ",drrive.ID)
	fmt.Println("Name: ",drrive.Name)
	fmt.Println("CarNo: ",drrive.CarNo)

	customer:= Customer{
		ID :1,
		Name : "Mekuannet",
		Phone : "0931166196",
	}
	customer1:= Customer{
		ID :1,
		Name : "Mekuannet",
		Phone : "0931166196",
	}

	jsonData, err:= json.MarshalIndent(
		customer,
		"",
		" ",
	)

	jsonData1, err:= json.MarshalIndent(
		customer1,
		"",
		" ",
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println(string(jsonData1))

	// json to struct converssion UnMarshalling
	data:= []byte(`
		{
			"id": 3,
			"name": "Abebe",
			"phone": "0987654321"
		}
	`)


	var customar2 Customer

	err1:= json.Unmarshal(data, &customar2)
	if err1!= nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(customar2.ID)
	fmt.Println(customar2.Name)
	fmt.Println(customar2.Phone)
}