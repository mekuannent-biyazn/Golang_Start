package main

import (
	"fmt"
	"slices"
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

	// practicar exercise for array
	arry5:= [5]int {1,2,3,4,5}
	fmt.Println(arry5)

	for i:=0; i<5; i++{
		fmt.Println(arry5[i])
	}

	// printing student names with loop in array
	studName:= [5]string {"Meku", "Abebe", "Kebede", "Alemu", "Zemenu"}
	for i:=0; i<5; i++{
		fmt.Println(studName[i])
	}

	// calculate students score
	scores := [5]int{70,80,90,85,95}
	sum:= 0
	for i:=0; i<5; i++{
		sum +=scores[i]
	}
	total:= sum
	avg:= total/len(scores)

	fmt.Println("Tatal score: ", total)
	fmt.Println("Average score: ",avg)

	// find the largest number
	numbers := [5]int{100,30,15,80,25}
	largest := numbers[0]

	for i:=1; i<5; i++{
		if numbers[i]>largest {
			largest= numbers[i]
		}
	}
	fmt.Println("The largest number is ", largest)

	addstudent()
}

func addstudent() {
	for i:=1; i<=5; i++{
		var name string
		fmt.Print("Enter new student Name: ")
		fmt.Scanln(&name)

		fmt.Println("The new student added is: ", name)
		
		var scors [5]int
		for j:=0; j<len(scors); j++{
			fmt.Printf("Enter score %d:", j+1)
			fmt.Scanln(&scors[j])
		}
		fmt.Println("Scors of student ",name,"is: ",scors)
		fmt.Println()
		total:=0
		var avg int

		for a:=0; a<len(scors); a++{
			total += scors[a]
			avg=total/len(scors)
		}
		fmt.Println("Total Scors of student ",name,"is: ",total)
		fmt.Println("Average Scors of student ",name,"is: ",avg)
	}
}

