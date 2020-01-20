package main

import (
	"fmt"
)

// ListNode for linkedlist
type ListNode struct {
	data string
	next *ListNode
}

func main() {
	linkedList := new(ListNode)
	fmt.Println(linkedList.data, linkedList.next)
	fmt.Printf("%T\b", linkedList)
}
