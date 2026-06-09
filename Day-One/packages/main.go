package main

import (
	"fmt"
	"packages/calculator"
	"strings"
	"strconv"
	"time"
)

func main(){
	sum:= calculator.Add(12, 21)
	fmt.Println("Sum: ",sum)

	sub:= calculator.Sub(12, 21)
	fmt.Println("Diffrence: ",sub)

	pro:= calculator.Mul(12, 21)
	fmt.Println("Product: ",pro)

	divs:= calculator.Div(12, 21)
	fmt.Println("Devision: ",divs)

	fmt.Println(strings.ToUpper("golang backend"))
	fmt.Println(strconv.Atoi("25"))
	fmt.Println(strings.Contains("golang","lang"))
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Second())
}