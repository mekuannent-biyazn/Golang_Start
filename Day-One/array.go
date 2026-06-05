package main

import (
	"fmt"
	// "slices"
)

func main() {
	var zeroArry [5]int
	var myArry = [3] int{10, 20, 30}
	var cars = [5] string{1: "volvo", 4: "BMW"}
	var snacks = [...] string {"chips", "cookies", "chocolate"}

	fmt.Println("Zero Array:", zeroArry)
	fmt.Println("My Array:", myArry)
	fmt.Println("Cars:", cars)
	fmt.Println("Snacks:", snacks)

	X := [3] int{10, 20, 30}
	y := [3] int{10, 20, 30}

	fmt.Println(len(snacks))
	fmt.Println(X == y)
	fmt.Println(X != y)

	X[0] = 100
	fmt.Println(X[0])
	fmt.Println(X)

	// slice
	someslice := [] int{1, 2, 3, 4, 5}
	var carslice = [] string {1: "volvo", 4: "BMW"}
	var sliceWithNovalue [] int

	fmt.Println(someslice)
	fmt.Println(carslice)
	fmt.Println(sliceWithNovalue)
	fmt.Println(sliceWithNovalue == nil)

	fruit1 := [] string {"orange", "banana", "avocado"}
	fruit2 := [] string {"cabage", "tomato", "potato", "hhhg"}

	fmt.Println(slices.Equal(fruit1,fruit2))
	fmt.Println(len(fruit1))
 	fruit1=append(fruit1, "Strobery")
	fruit1 = append(fruit1, fruit2...)
	fmt.Println(fruit1)
	fmt.Println(len(fruit1))

	var num [] int
	fmt.Println(num, len(num), cap(num))
	num = append(num, 1)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 2)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 3)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 4)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 5)
	fmt.Println(num, len(num), cap(num))

	num2:= make([] int, 0, 20)
	num2 = append(num2, 5, 10, 15, 20)
	fmt.Println(num2, len(num2), cap(num2))

	num3 := []int {3, 4, 6, 8, 12}
	fmt.Println(num3 [1:4])
	fmt.Println(num3 [1:])
	fmt.Println(num3 [:4])
	fmt.Println(num3 [:])

	num4 := num3[1:4]
	num4[1]= 1994
	fmt.Println(num4[:])
	fmt.Println(num3)

	s:= []int {10,20,30,40}
	s= append(s[:2], s[2+1:]...)
	fmt.Println(s)

	values := []int {4,5,6,7}
	newValues := make([]int, 4)
	copy(newValues, values)
	newValues[0]= 1994
	fmt.Println(newValues)
	fmt.Println(values)
}