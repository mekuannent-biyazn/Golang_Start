package main
import "fmt"

// range in go with array
func main() {
	fmt.Println("recursive function!")

	fmt.Println(factorial(5))
	fmt.Println("Fibonacy sequence of 10 :",fibonacy(10))
}

// factorial
func factorial(n int) int{
	if (n==0){
		return 1
	}else{
		return n*factorial(n-1)
	}
}

// fibonaci sequence 
func fibonacy(n int) int{
	if (n==0){
		return 1
	}else if(n==1){
		return 1
	}else {
		return n+fibonacy(n-1)
	}
}