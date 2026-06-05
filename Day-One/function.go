package main

import "fmt"

// function with parameters and return value
func add(a int, b int) int {
	return a+b
}

// function with multiple return values
func calc(a, b int) (int, int) {
	return a+b, a*b
}

// employ saalary calculation
func salary(slry float64) float64 {
	if slry > 20000{
		slry= slry*0.2 + slry
		// fmt.Printf("Your salary is %.2f\n", slry)
	} else if slry >=10000 {
		slry = slry*0.1 + slry
		// fmt.Printf("Your salary is %.2f\n", slry)
	} else {
		slry = slry*0.05 + slry
		// fmt.Printf("Your salary is %.2f\n", slry)
	}
	return slry
}

// Student Result System using Functions.
func getGrade(score float64) string {
	if score >= 90 { 
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func isPased(score float64) bool {
	return score >= 60
}

func displayResult(name string, score float64) {
	grade := getGrade(score)
	passed := isPased(score)

	fmt.Printf("Student: %s\n", name)
	fmt.Printf("Score: %.2f\n", score)
	fmt.Printf("Grade: %s\n", grade)
	fmt.Printf("Passed: %t\n", passed)
}

func main(){
	fmt.Println("The sum of 5 and 3 is:", add(5,3))
	sum, product := calc(9,5)
	fmt.Printf("The sum of 9 and 5 is %d and their product is %d\n", sum, product)	

	var slryInput float64
	fmt.Print("Enter your salary: ")
	fmt.Scanln(&slryInput)

	slary:= salary(slryInput)
	fmt.Printf("Your salary is %.2f\n", slary)

	// // functions with  no return value
	// sum, _ := calc(10, 5)

	// fmt.Println(sum)

	// Student Result System
	var name string
	var score float64 

	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter student score: ")
	fmt.Scanln(&score)

	displayResult(name, score)
}