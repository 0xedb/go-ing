package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 100
	fmt.Printf("%v \n", <-ch)

}
