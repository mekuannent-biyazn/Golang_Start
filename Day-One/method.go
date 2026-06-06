package main

import "fmt"

type User struct{
	Name string
	Age int
}

type Techer struct{
	Id int
	Name string
	EduL string
}

type Employee struct{
	Name string
	Salary float64
}

type Rectangle struct{
	Width float64
	Hieght float64
}

// Method Returning Values
func (r Rectangle) area() float64{
	return r.Width * r.Hieght
}

// Method with parameters.
func (e Employee) bonus(percent float64){
	salary:= e.Salary * percent/100 + e.Salary

	fmt.Println("Your salary with bonus is : ",salary)
}

// Method to disply TEACHER
func (t Techer) show(){
	fmt.Println("Teacher Id        : ",t.Id)
	fmt.Println("Teacher Name      : ",t.Name)
	fmt.Println("Teacher Edu.Leble : ",t.EduL)
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

	teacher:= Techer{
		Id: 1,
		Name: "Mekashaw",
		EduL: "BSC",
	}

	employee:= Employee{
		Name: "Mekuannet",
		Salary: 45000,
	}

	rectangle:= Rectangle{
		Width: 23.6,
		Hieght: 12.4,
	}

	user.desply()
	fmt.Println("===========Teacher Info============")
	teacher.show()
	fmt.Println("===========Employee Info============")
	employee.bonus(20.0)
	fmt.Println("===========Area of Rectangle============")
	Area:= rectangle.area()
	fmt.Println(Area)
}