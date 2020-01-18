package main

import (
	"fmt"
	"sync"
)

func conCur(msg string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for st := 1; st <= 1000; {
			fmt.Println(msg, "-----", st)
			st++
		}
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	conCur("bruno", &wg)
	conCur("Kofi", &wg)
	wg.Wait()
}
