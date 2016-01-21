package main

import(
	"fmt"
	"time"
)


func Producer(queue chan<-int){
	
	for i := 0;i<10;i++{
			fmt.Println("produce :",i)
			queue <- i
		

	}
}


func Consumer(queue <- chan int){
	for i := 0;i<10;i++{
	
		v := <-queue
		fmt.Println("consume: ",v)
	
	}

	
}


func main(){
	queue := make(chan int,2)
	go Producer(queue)
	go Consumer(queue)
	fmt.Println("start")
	time.Sleep(1e9)

}