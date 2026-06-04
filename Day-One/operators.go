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
}