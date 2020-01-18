package concurrency

import "fmt"

func Receiver(ch chan int) { 
		fmt.Printf("Received: %v", <-ch)
}

func Sender(ch chan int) { 
		ch <- 90
		fmt.Printf("Sent: %v", ch)
 
}
