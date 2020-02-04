package main

import "fmt"

func putter(w int, ch *chan int) {
	*ch <- w
	close(*ch)
}

func main() {
	ch := make(chan int)
	go putter(27800, &ch)

	for i := range ch {
		fmt.Println("-----", i)
	}
}
