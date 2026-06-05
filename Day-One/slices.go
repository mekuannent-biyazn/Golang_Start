package main

import "fmt"

func main(){
	var number []int
	number= append (number, 10)
	fmt.Println(number)

	slis:= []int{}
	slis= append(slis, 20, 30, 40)

	number= append(number, slis...)
	fmt.Println(number)
	fmt.Println(slis)

	// slice loop
	for i:=0; i<len(slis); i++{
		fmt.Println(slis[i])
	}

	// range slice loop
	for index, value:= range number{
		fmt.Println(index, value)
	}

	// ignor index 
	for _, value:= range number{
		fmt.Println(value)
	}

	// Slice Internals
	fmt.Println(len(number))
	fmt.Println(cap(number))
}