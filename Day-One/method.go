package main

import "fmt"

type User struct{
	Name string
	Age int
	Email string
}

func (u User) IsAdmin() bool{
	return u.Email == "Admin@gmail.com"
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

	// =============================================

	type bankAccount struct{
		owner string
		Balance float64
	}

	func (ba *bankAccount) deposit( amount float64 ) {
		fmt.Println("===========Deposite==========")
		ba.Balance += amount 
		fmt.Println("Owner: ", ba.owner,"")
		fmt.Println("Balance : ", ba.Balance)
		fmt.Println("Diposite: ", amount)
		fmt.Println("Balance : ", ba.Balance)
	} 

	func (ba *bankAccount) Withdraw(amount float64){
		fmt.Println("===========Withdrwal==========")
		ba.Balance -= amount 
		fmt.Println("Owner: ", ba.owner,"")
		fmt.Println("Balance : ", ba.Balance)
		fmt.Println("Withdrwal: ", amount)
		fmt.Println("Balance : ", ba.Balance)
	}

	func (ba bankAccount) disply(){
		fmt.Println("=============Disply the result================")
		fmt.Println("Owner : ",ba.owner)
		fmt.Println("Balance: ",ba.Balance)
	}

func main(){
	greet()
	user:= User{
		Name: "MEKUANNENT",
		Age: 25,
		Email: "Admins@gmail.com",
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
	fmt.Println("===========Is Admin============")
	admin:= user.IsAdmin()
	fmt.Println(admin)
	if user.IsAdmin() {
		fmt.Println("Allow admin access")
	} else{
		fmt.Println("Unauthorized access!!")
	}

	baAct:= bankAccount{
		owner: "Mekuannent",
		Balance: 3000,
	}
	baAct.disply()
	baAct.deposit(1000.0)
	baAct.Withdraw(500.0)
}