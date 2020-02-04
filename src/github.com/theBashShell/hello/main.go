package main


import "fmt"

type str string

/*Human is a human prototype */
type Human struct {
	name string
	age uint16
	isMarried bool
}

func (h *Human) info() string {
	return fmt.Sprintf("%s----", h.name)
}

func (st str) more() {
	fmt.Println(st)
}

func main() {
	bruno := new(Human)
	bruno.name = "bruno kofi edoh"
	var ss str = "bru!"
	var myAge uint16 = 200
	fmt.Println(myAge)
	fmt.Println(bruno)
	fmt.Println(bruno.info())
	ss.more()
}