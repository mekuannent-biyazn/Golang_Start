package main

import (
	"fmt"
	"time"
	"sync"
)

func task1(){
	fmt.Println("Task1 started.")
	fmt.Println("waiting for 2 Seconds...\n\n")
	time.Sleep(2*time.Second)
	fmt.Println("Task1 finished.\n\n")
}

func task2(){
	fmt.Println("Task2 started.")
	fmt.Println("waiting for 2 Seconds...\n\n")
	time.Sleep(2*time.Second)
	fmt.Println("Task2 finished.\n\n")
}

func Hello(){
	fmt.Println("Hollo from Goroutin!!")
}

func PrintNumbers(){
	for i:=1; i<=5; i++{
		fmt.Println("Number: ",i)
		time.Sleep(500* time.Millisecond)
	}
}

func PrintLetters(){
	for _,ch:= range ("ABCDE"){
		fmt.Println("Letter: ",string(ch))
		time.Sleep(500* time.Millisecond)
	}
}

func task(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Task is running...")
}

func MulWaitGroup(){
	var wg sync.WaitGroup
	for i:=1; i<=5; i++ {
		wg.Add(1)
		go func (num int){
			defer wg.Done()
			fmt.Println("Number", num)
		}(i)
	}
	wg.Wait()
}

func main(){
	task1()
	task2()

	go Hello()

	go PrintNumbers()
	go PrintLetters()
	time.Sleep(4* time.Second) // to protect the main finction exit before hello function

	var wg sync.WaitGroup

	wg.Add(1)
	go task(&wg)
	wg.Wait()
	fmt.Println("Process Finished.")

	MulWaitGroup()
}