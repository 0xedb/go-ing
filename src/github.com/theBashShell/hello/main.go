package main

import "fmt"

func putter(w int, ch *chan int) {
	*ch <- w
}

func main() {
	ch := make(chan int)
	go putter(200, &ch)

	w := <-ch
	fmt.Println(w)
}
