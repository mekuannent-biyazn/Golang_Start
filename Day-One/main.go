package main
import (
	"fmt"
	"time"
)

func main(){
	currentYear := time.Now().Year()
	fmt.Println("Hello, World!");
	fmt.Println("My name is Mekuannent Biyazn")
	fmt.Println("current year is", currentYear)
	fmt.Println("I am ", currentYear -2001 ," years old")
}