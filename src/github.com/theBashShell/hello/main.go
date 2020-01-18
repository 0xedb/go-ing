package main

import "fmt"


func doubleCallback(num int, callback func(n int) string) string{
	return callback(num)
}
 
func main() {
	echo := func(n int) string {
	return fmt.Sprintf("%v-----", n)
}
	fmt.Println(doubleCallback(10, echo))
}
