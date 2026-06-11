package main

import (
	"fmt"
	"context"
	"time"
)

func worker(ctx context.Context){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("Work is Done.")
			return
		default :
		fmt.Println("Working...")
		time.Sleep(time.Second)
		}
	}
}

func verify( cxt context.Context){
	for{
		select{
		case <-cxt.Done():
			fmt.Println("Click here to get the new link.")
			return
		default :
		time.Sleep(3*time.Second)
		fmt.Println("This link is valid for 3 seconds!")
		}
	}
}

func main(){
	cxt:= context.Background()
	fmt.Println(cxt)

	// context withCancel
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)
	time.Sleep(5*time.Second)
	cancel()
	time.Sleep(time.Second)

	// context Withtimeout
	ctxt, cancel:= context.WithTimeout(context.Background(), 3*time.Second)
	go verify(ctxt)
	time.Sleep(5*time.Second)

	defer cancel()
	<-ctxt.Done()

	fmt.Println("Timeout reached!")
}