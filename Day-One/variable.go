package main

import "fmt"

func main(){
	// Variable declaration and initialization

	var name string ="Mekuannent Biyazn"
	MyName:= "Mekuannent Biyazn"
	Age:=25
	var isStudent bool = false
	var educationLevel string
	fmt.Println(name)
	fmt.Println("My name is",MyName)
	fmt.Println("I am ",Age," Years old.")
	fmt.Println("Am I a student?",isStudent)
	educationLevel = "Bachelor's Degree in Computer Science"
	fmt.Println("My education level is",educationLevel)

	fmt.Printf("My education level is %v\n", educationLevel)

	var grade int = 375
	fmt.Printf("My grade is %d\n", grade)
	fmt.Println("My grade is ", grade)

	// constants
	const pi float64 = 3.14159
	const gravity float64 = 9.81
	fmt.Println("The value of pi is", pi)
	fmt.Println("The value of gravity is", gravity)

	// recived input from user
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)
	fmt.Println("Hello", userName, "Welcome to Go programming!")

	// simple area calculation for a circle
		var radius float64
		fmt.Print("Enter the radius of the circle: ")
		fmt.Scanln(&radius)
		area := pi * radius * radius 
		fmt.Println("The area of the circle is", area)

	// student information
	var studentName string
	var studentAge int

	fmt.Print("Enter student name: ")
	fmt.Scanln(&studentName)

	fmt.Print("Enter student age: ")
	fmt.Scanln(&studentAge)

	fmt.Printf("Student Name: %s, Student Age: %d\n", studentName, studentAge)

	// multiple variable declaration and simple arithmetic operation
	var a, b, sum, sub, mul, div int
	fmt.Print("Enter first and second number: ")
	fmt.Scanln(&a, &b)
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)
	fmt.Printf("The difference of %d and %d is %d\n", a, b, sub)
	fmt.Printf("The product of %d and %d is %d\n", a, b, mul)
	fmt.Printf("The quotient of %d and %d is %d\n", a, b, div)

	// advanced printing with formatting
	var name2 string
	var age2 int
	var University string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name2)
	fmt.Print("Enter your age  : ")
	fmt.Scanln(&age2)
	fmt.Print("Enter your university : ")
	fmt.Scanln(&University)

	fmt.Printf("Your name is %s, Your age is %d and you are studying at %s\n", name2, age2, University)

	//type conversion
	var num1 int = 10
	var num2 float64 = 5.5
	
	sum2 := float64(num1) + num2
	fmt.Printf("the sum of %d and %f is %f\n", num1, num2, sum2)
	sum := num1 + int(num2)
	fmt.Printf("the sum of %d and %f is %d\n", num1, num2, sum)
}