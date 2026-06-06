package main

import "fmt"

type student struct{
		name string
		age int
	}

type Employee struct{
	Id int
	name string
	position string
	salary int
}

func displyStd(stud student){
	fmt.Println(stud.name)
	fmt.Println(stud.age)
}

func main(){
		st1:= student{
		name: "mekuannent",
		age: 25,
	}

	fmt.Println("Name: ", st1.name)
	fmt.Println("Age: ", st1.age)

	st1.age= 26
	fmt.Println("Age after update: ", st1.age)

	var STD student
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scanln(&n)

	for i:=1; i<=n; i++{
		fmt.Printf("Enter student %d Name: ", i)
		fmt.Scanln(&STD.name)
		fmt.Printf("Enter student %d Age: ", i)
		fmt.Scanln(&STD.age)

		fmt.Printf("Student %d Name is %s\n", i, STD.name)
		fmt.Printf("Student %d Age is %d\n", i, STD.age)
		fmt.Println()
	}

	// multiple fields
	emp:= Employee{
		Id : 1,
		name: "Mekuannent",
		position: "developer",
		salary: 25000,
	}
	fmt.Println("Employee Id        : ", emp.Id)
	fmt.Println("Employee Name      : ", emp.name)
	fmt.Println("Employee Position  : ", emp.position)
	fmt.Println("Employee Salary    : ", emp.salary, "\n")

	// struct + function
	displyStd(st1)
}