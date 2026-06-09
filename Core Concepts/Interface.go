package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak(){
	fmt.Println("Woof")
}

type Cat struct{}

func (c Cat) Speak(){
	fmt.Println("Maow")
}

func main(){
	var s Speaker

	s= Dog{}
	s.Speak()

	s= Cat{}
	s.Speak()
}