package main

import (
	"fmt"
	"time"
	"sync"
)

func sendMessage(ch chan string){
	time.Sleep(2*time.Second)
	ch<-"Raide Accepted."
}

func SendMsgToCustomer(ch chan string){
	time.Sleep(2*time.Second)
	ch<-"The driver is Arrived. Start comming to car station."
}

// func send(ch chan<- int){
// 	ch<-100
// }

// func recive(ch <-chan int){
// 	time.Sleep(1*time.Second)
// 	go send(ch)

// 	val:= <-ch
// 	fmt.Println("Channel: ",val)
// }

func processRequest(Id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Theprocess request %d\n", Id)
}

func featchData(ch chan string){
	result:="The data is featched successfully!."
	ch<- result 
}

func Unbuffered(ch chan string){
    ch<-"Unbuffered message"
}

func driver(requests chan string){
	for {
		request:=<-requests
		fmt.Println("Request is accepted:",request)
	}
}

func Produce(ch chan int){
	for i:=0; i<5; i++{
		ch<-i
	}
	close(ch)
}

func main(){
	var wg sync.WaitGroup

	prd:=make(chan int, 5)
	go Produce(prd)

	for val:= range prd{
		fmt.Printf("Product %d\n",val)
	}
	// what hapen when buffur is full
	bff:=make(chan int, 2)
	bff<-10
	bff<-20
	fmt.Println("The buffer is full!")
	time.Sleep(2*time.Second)
	close(bff)

	// ubfr:=make(chan int)
	// ubfr<-40
	// fmt.Println("Waiting Reciever.")
	// recr:=<-ubfr
	// fmt.Printf("%d \nAfter this Deadlock Occured.",recr)
	// // ubfr<-50
	// time.Sleep(2*time.Second)
	// close(ubfr)

	// ride requesting Sym
	raidRequest:= make(chan string)

	// wg.Add(1)
		// defer wg.Done()
	go driver(raidRequest)

	raidRequest<-"Customer A"
	raidRequest<-"Customer B"
	raidRequest<-"Customer C"

	time.Sleep(5*time.Second)

	
	// wg.Wait()

	// Unbuffered channel
	ubf:=make(chan string)
	for i:=1;i<=5;i++{
    	wg.Add(1)
		go func(id int){
			defer wg.Done()
			Unbuffered(ubf)
		}(i)
		msg:=<-ubf
		fmt.Printf("Data: %s\n",msg)
	}
	wg.Wait()

	bfr:=make(chan string, 3)
	wg.Add(1)
	go func(){
		defer wg.Done()
		bfr<-"Piza"
		bfr<-"Burger"
		bfr<-"Delcious food"
		close(bfr)

		for order:= range bfr{
			fmt.Printf("Ordered Food: %s\n", order)
		}

		book:=make(chan string, 2)

		book<- "Book1"
		book<- "Book2"

		fmt.Println(<-book)
		fmt.Println(<-book)
	}()
	wg.Wait()
	
	ch:=make(chan int)

	go func(){
		ch <- 100
	}()

	for i:=1; i<=5; i++{
		wg.Add(1)
		go func(id int){
			processRequest(id, &wg)
		}(i)
	}
	wg.Wait()

	chfd:=make(chan string)
	wg.Add(1)
	go func(){
		defer wg.Done()
		featchData(chfd)
	}()
	data:=<-chfd
	fmt.Printf("The featch data: %s\n",data)
	wg.Wait()
	// time.Sleep(1*time.Second)

	val:= <-ch
	fmt.Println("Channel: ",val)

	ch2:=make(chan string)
	go sendMessage(ch2)

	vall:=<-ch2
	fmt.Println(vall)

	ch3:=make(chan string)
	go SendMsgToCustomer(ch3)

	smsc:=<-ch3
	fmt.Println(smsc)

	// Buffered channel
	num:=make(chan int, 3)

	num<-10
	num<-20
	num<-30

	fmt.Println(<-num)
	fmt.Println(<-num)
	fmt.Println(<-num)

	// chn:=make(chan int)
	// go recive(chn)

	chn:=make(chan int)
	go func(){
		chn<-100
	}()
	value:=<-chn
	close(chn)
	fmt.Println(value)
	values, ok := <-chn
	if !ok{
		fmt.Println("The Cannel is not closed.",values)
	} else {
		fmt.Println("The chanell is closed.")
	}

	go func (){
	for i:=1; i<=5; i++{
		ch<-i
	}
	close(ch)
}()
for num:= range (ch){
	fmt.Println(num)
}
}