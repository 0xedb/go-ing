package main

import (
	"fmt"
	"time"
)

func printer(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	go printer("first")
	printer("second")

}
