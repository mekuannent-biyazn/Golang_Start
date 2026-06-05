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

	Mksl:= make([]int, 0, 20)
	Mksl= append(Mksl, 10, 20, 30)

	fmt.Println(Mksl, len(Mksl), cap(Mksl))

	// coping slices
	src:= []int {1,2,3}
	dest:= make([]int, len(src))
	copy(dest, src)
	dest = append(dest, 4,5,6)
	fmt.Println("The source slice is : ",src)
	fmt.Println("The destination slice is : ",dest)

	prts:= []int {10, 20, 30, 40, 50}
	for _, value:= range prts{
		fmt.Println(value)
	}

	var n int
	fmt.Print("Enter required number of students: ")
	fmt.Scanln(&n)
	Stname:= [] string {}
	var name string
	for i:=0; i<=n; i++{
		fmt.Printf("Enter student %d ", i+1)
		fmt.Scanln(&name)
		Stname= append(Stname, name)
	}
	fmt.Println(Stname)
	fmt.Println(Stname[1:4])
}