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

	// Mini Projects
	type StudentMinP struct {
		Id int
		name string
		dpt string
		GPA float64
	}

	// create 3 students
	Mpstud1:= StudentMinP{
		Id: 1,
		name: "Mekuannent",
		dpt: "CS",
		GPA: 3.71,
	}
	
	Mpstud2:= StudentMinP{
		Id: 2,
		name: "Maritu",
		dpt: "IT",
		GPA: 3.8,
	}

	Mpstud3:= StudentMinP{
		Id: 3,
		name: "Tiruya",
		dpt: "Accounting",
		GPA: 3.7,
	}

	fmt.Println("Student 1 info: \nId: \t",Mpstud1.Id,"\nName: \t",Mpstud1.name,"\nDepartment: \t",Mpstud1.dpt,"\nCGPA: \t",Mpstud1.GPA)
	fmt.Println("\n\nStudent 2 info: \nId: \t",Mpstud2.Id,"\nName: \t",Mpstud2.name,"\nDepartment: \t",Mpstud2.dpt,"\nCGPA: \t",Mpstud2.GPA)
	fmt.Println("\n\nStudent 3 info: \nId: \t",Mpstud3.Id,"\nName: \t",Mpstud3.name,"\nDepartment: \t",Mpstud3.dpt,"\nCGPA: \t",Mpstud3.GPA)
		
	// store students in slice
	SStudents:= []StudentMinP{
		{
			Id: 1,
			name: "Mekuannent",
			dpt: "CS",
			GPA: 3.71,
		},
		{
			Id: 2,
			name: "Maritu",
			dpt: "IT",
			GPA: 3.8,
		},
		{
			Id: 3,
			name: "Tiruya",
			dpt: "Accounting",
			GPA: 3.7,
		},
	}
	for _, students:= range SStudents{
		fmt.Println("\n==================================")
		fmt.Println("Id: ", students.Id)
		fmt.Println("Name: ", students.name)
		fmt.Println("Department: ", students.dpt)
		fmt.Println("GPA: ", students.GPA)
	}
}