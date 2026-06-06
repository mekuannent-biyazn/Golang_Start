package main

import "fmt"

type User struct{
	Name string
	Age int
}

// Normal functio
	func greet(){
		fmt.Println("Hello every one!!")
	}

	// method
	func (u User) desply(){
		fmt.Println("Name: ",u.Name)
		fmt.Println("Age: ",u.Age)
	}

func main(){
	greet()
	user:= User{
		Name: "MEKUANNENT",
		Age: 25,
	}
	user.desply()
}