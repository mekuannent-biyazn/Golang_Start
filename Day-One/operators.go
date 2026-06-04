package main
import "fmt"

func main(){
	// Arithmetic operators
	a:= 10
	b:= 4

	fmt.Println("addition:", a+b)
	fmt.Println("subtraction:", a-b)
	fmt.Println("multiplication:", a*b)
	fmt.Println("division:", a/b)
	fmt.Println("modulus/remainder:", a%b)

	//comparison operators
	fmt.Println("equal to:", a==b)
	fmt.Println("not equal to:", a!=b)
	fmt.Println("greater than:", a>b)
	fmt.Println("less than:", a<b)
	fmt.Println("greater than or equal to:", a>=b)
	fmt.Println("less than or equal to:", a<=b)

	//logical operators
	fmt.Println("Logical AND:", (a>5) && (b<5))
	fmt.Println("Logical OR:", (a>5) || (b>5))
	fmt.Println("Logical NOT:", !(a>5))

	fmt.Println("Logical AND:", (a<5) && (b<5))
	fmt.Println("Logical OR:", (a<5) || (b>5))
	fmt.Println("Logical NOT:", !(a<5))

	// if statement
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age <= 18 {
		fmt.Println("You are a minor.")
	}

	// if-else statement
	if age <= 18 {
		fmt.Println("You are a child.")
	}	else if age > 18 && age < 50 {
		fmt.Println("You are an adult.")
	}else{
		fmt.Println("You are a senior citizen and you are old.")
	}

	// if-else if-else statement
	var gpa float64
	fmt.Print("Enter your GPA: ")
	fmt.Scanln(&gpa)

	if gpa >= 3.5 {
		fmt.Println("You are an excellent student.")
	} else if gpa >= 2.0 && gpa < 3.5 {
		fmt.Println("You are a good student.")
	} else {
		fmt.Println("You need to work harder.")
	}

	// switch statement
	var day string
	fmt.Print("Enter a day of the week: ")
	fmt.Scanln(&day)

	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	case "Friday":
		fmt.Println("Today is Friday.")	
	case "Saturday":
		fmt.Println("Today is Saturday.")
	case "Sunday":
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Invalid day of the week.")
	}

	// login validation
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	if username == "admin" && password == "password" {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Invalid username or password.")
	}

	// user acess level
	role := "admin"

	switch role {
	case "admin":
		fmt.Println("you have full access")
	case "editor":
		fmt.Println("you have edit access")
	case "viewer":
		fmt.Println("you have view access")
	default:
		fmt.Println("invalid role")
	}

	// student grade evaluation
	var grade float64
	fmt.Print("Enter your grade: ")
	fmt.Scanln(&grade)

	if grade >= 90 && grade <= 100 {
		fmt.Println("You got an A.")
	} else if grade >= 80 && grade < 90 {
		fmt.Println("You got a B.")
	} else if grade >= 70 && grade < 80 {
		fmt.Println("You got a C.")
	} else if grade >= 60 && grade < 70 {
		fmt.Println("You got a D.")
	} else {
		fmt.Println("You got an F.")
	}

	// Employee Bonus System
	var salary float64
	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)

	if salary >= 20000{
		salary =salary*0.2 + salary
		fmt.Println("You get a 20% bonus. Your new salary is", salary)
	} else if salary >= 10000 && salary < 20000 {
		salary = salary*0.1 + salary
		fmt.Println("You get a 10% bonus. Your new salary is", salary)
	} else {
		salary = salary*0.05 + salary
		fmt.Println("You get a 5% bonus. Your new salary is", salary)
	}

	// Student Result Management Program:
	var Sname string
	var Score float64
	
	fmt.Print("Enter student name: ")
	fmt.Scanln(&Sname)
	fmt.Print("Enter student score: ")
	fmt.Scanln(&Score)

	if Score >= 90 && Score <= 100 {
		fmt.Println("Student Name: ",Sname, "\nScore :", Score, "\nGrade : A", "\nstatus: Pass")
	} else if Score >= 80 && Score < 90 {
		fmt.Println("Student Name: ",Sname, "\nScore :", Score, "\nGrade : B", "\nstatus: Pass")
	} else if Score >= 70 && Score < 80 {
		fmt.Println("Student Name: ",Sname, "\nScore :", Score, "\nGrade : C", "\nstatus: Pass")
	} else if Score >= 60 && Score < 70 {
		fmt.Println("Student Name: ",Sname, "\nScore :", Score, "\nGrade : D", "\nstatus: Pass")
	} else {
		fmt.Println("Student Name: ",Sname, "\nScore :", Score, "\nGrade : F", "\nstatus: Fail")
	}
}