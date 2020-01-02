package main

import "fmt"

func init() {
	fmt.Println("Doing this first")
}

func init() {
	fmt.Println("Doing this second")
}

func one() int {
	return 1
}

func zero() int {
	return 60
}

func main() {
	w := 'K'
	fmt.Printf("%T\n\n", w)
	fmt.Println(one())
	fmt.Println(zero())
}
