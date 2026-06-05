package main

import "fmt"

func main() {
	// create a map
	student:= map[string]string{
		"name": "mekuannent",
		"Department": "Computer Sicience",
	}
	fmt.Println(student["name"])
	fmt.Println(student["Department"])

	// adding new data in to map
	addstd:= map[string]string {}
	addstd["name"]="Mekuannent"
	addstd["city"]="Addis ababa"

	addstd["city"]="Bahir Dar"
	delete(addstd, "city")

	value, exists:= addstd["city"]
	fmt.Println(value)
	fmt.Println(exists)

	value1, exists:= addstd["name"]
	fmt.Println(value1)
	fmt.Println(exists)

	fmt.Println(addstd["name"])
	fmt.Println(addstd["city"])

	// Loops in map
	fmt.Println("===============loop with map================")

	Student:= map[string] string{
		"Name": "Mekuannent",
		"Country": "Ethiopia",
		"City": "Addis Ababa",
	}
	for key, value:= range Student{
		fmt.Println(key, value)
	}

	// Map with Different Types
	result:= map[string] int{
		"Maath": 90,
		"Physics": 85,
	}
	for key, value:= range result{
		fmt.Println(key, value)
	}

	user:= map[string] string{
		"admin": "123",
		"johne": "abc",
	}
	username:= "admin"
	password:= "123"

	if user[username]== password{
		fmt.Println("Login sucessFully!!")
	} else {
		fmt.Println("Login Failed!!")
	}
}