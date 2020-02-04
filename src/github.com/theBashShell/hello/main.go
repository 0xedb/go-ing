package main

import (
	"fmt"
	// "time"
)

func putter(ch *chan int) {
	*ch <- 1
	close(*ch)
}

func another(ch *chan int) {
	*ch <- 2
	// time.Sleep(time.Millisecond * 100)
	close(*ch)
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go putter(&ch)
	go another(&ch2)

	select {
	case one := <-ch:
		fmt.Println("channel 1: ", one)
	case two := <-ch2:
		fmt.Println("channel 2: ", two)
	}

}
