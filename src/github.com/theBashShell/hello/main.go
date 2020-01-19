package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup, slice *[]int) {
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i 
		}
		close(ch)
		defer wg.Done()
	}()
	defer fmt.Println(*slice)
}

func consumer(ch chan int, wg *sync.WaitGroup, mut *sync.Mutex) {
	wg.Add(1)
	go func() {
		for i := range ch {
			fmt.Printf("***\tconsumed:\t%v\n", i)
		}
		defer wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	var slice []int

	channel := make(chan int)
	producer(channel, &wg, &slice) 
	wg.Wait()
}
