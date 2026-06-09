package main

import (
	"fmt"
	"math"
)

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

// Payment methods
type PaymentMethod interface{
	Pay(ammount float64)
}

type TellBirr struct{}

func (tb TellBirr) Pay(amount float64){
	fmt.Println("The payment in TelleBirr with amount: ",amount)
}

type Cash struct{}

func (c Cash) Pay(amount float64){
	fmt.Println("The cash payment is completed with amount: ", amount)
}

func processPayment(p PaymentMethod){
	p.Pay(1000)
}

// Calculating Areas
type Shape interface {
	Area() float64
}

type Measurable interface{
	Peremeter() float64
}

type Geometry interface{
	Shape
	Measurable
}

type Rectangle struct{
	width, hieght float64
}

func (r Rectangle) Area() float64{
	return r.width* r.hieght
}

func (r Rectangle) Peremeter() float64{
	return 2*(r.width* r.hieght)
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi* c.radius* c.radius
}

func (c Circle) Peremeter() float64{
	return 2*(math.Pi* c.radius* c.radius)
}

func CalculateArea(s Shape) float64{
	return s.Area()
}

func describeShape(g Geometry){
	fmt.Println("Area: ",g.Area())
	fmt.Println("Perimetre: ",g.Peremeter())
}

func main(){
	var s Speaker

	s= Dog{}
	s.Speak()

	s= Cat{}
	s.Speak()

	processPayment(Cash{})
	processPayment(TellBirr{})

	rect:= Rectangle{width: 5, hieght:7,}
	circle:= Circle{radius: 6}

	fmt.Println("Rectangle Area: ", CalculateArea(rect))
	fmt.Println("Circle Area: ", CalculateArea(circle))

	masterBox:= interface{}("type")
	describeValue(masterBox)

	checkInt, ok := masterBox.(int)
	if ok{
		fmt.Println("The Type is intger", checkInt)
	} else {
		fmt.Println("The type is not intger!")
	}

	fmt.Println("For Rectangle")
	describeShape(rect)

	fmt.Println("For Circle")
	describeShape(circle)
}

func describeValue(t interface{}){
	fmt.Printf("Type: %T, Value: %v\n",t,t)
}