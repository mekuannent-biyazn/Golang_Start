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

func main(){
	fmt.Println("The sum of 5 and 3 is:", add(5,3))
	sum, product := calc(9,5)
	fmt.Printf("The sum of 9 and 5 is %d and their product is %d\n", sum, product)	
}