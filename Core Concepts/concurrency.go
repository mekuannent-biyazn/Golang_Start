package main

import (
	"fmt"
	"time"
    "math/rand"
	"sync"
)

type Order struct{
	Id int
	Status string
	mu sync.Mutex
}

var(
	totalUpdate int
	updateMutex sync.Mutex
)

func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	orders:= generateOrders(20)

	// go func(){
	// 	processOrders(orders)
	// 	defer wg.Done()
	// }()

	// go func(){
	// 	updateOrderStatus(orders)
	// 	defer wg.Done()
	// 	}()

	// go func(){
	// 	reportOrderStatus(orders)
	// 	defer wg.Done()
	// }()

	for i:=0; i<3; i++{
		go func(){
			defer wg.Done()
			for _, order:= range orders{
				updateOrderStatus(order)
			}
		}()
	}

	wg.Wait()
	fmt.Println("\nThe process is finished. Exiting...")
	fmt.Println(totalUpdate)
}

func updateOrderStatus(order *Order){
	// for _, order:= range order{
		time.Sleep(
			time.Duration(50)*time.Millisecond,
		)
		status:= []string{
			"Pending", "Shiped", "Delivered",
		}[rand.Intn(3)]
		order.Status= status
		fmt.Printf("Update order %d Status %s\n", order.Id,status,)
	// }
	curruntUpdate:= totalUpdate
	time.Sleep(5* time.Millisecond)
	totalUpdate:=curruntUpdate + 1
}

func processOrders(orders []*Order){
	for _, order:= range orders{
		time.Sleep(
			time.Duration(50)*time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.Id)
	}
}

func generateOrders(count int) []*Order{
	orders:=make([]*Order, count)
	for i:=0; i<count; i++{
		orders[i]= &Order{
			Id: i, Status: "Pending",
		}
	}
	return orders
}

func reportOrderStatus(orders []*Order){
	for i:=0; i<5; i++{
		time.Sleep(1*time.Second)
		fmt.Println("\n-----Order Status------")
		for _,order:= range orders{
			fmt.Printf("Order %d Status %s\n",order.Id, order.Status,)
		}
		fmt.Println("-----------------------------\n")
	}
}