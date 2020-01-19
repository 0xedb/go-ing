package main

import (
	"fmt"
	"sync"
)

func producer(ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {

		fmt.Println("producer")
		ch <- "producer"

		defer wg.Done()
	}()
}

func consumer(ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {

		received := <-ch
		fmt.Println("consumed:\t", received)

		defer wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan string)
	producer(channel, &wg)
	consumer(channel, &wg)
	wg.Wait()
}
