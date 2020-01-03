package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
	"github.com/theBashShell/beautiful"
)

type Human struct {
	name string
	age uint16
}

type Walk interface {
	Walking()
}

func (h Human) Walk() {
	fmt.Printf("%s is walking...", h.name)
}

func main() {
	fmt.Println("main her!!!!!e")
	fmt.Println(beautiful.Doubler(339))
	fmt.Println(stringutil.Reverse("love"))

	bruno := Human{"bruno", 1}
	fmt.Println(bruno)
	bruno.Walk()

}
