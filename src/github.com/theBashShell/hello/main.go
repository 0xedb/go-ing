package main


import "fmt"

/*Human is a human prototype */
type Human struct {
	name string
	age uint16
	isMarried bool
}

func (h *Human) info() string {
	return fmt.Sprintf("%s----", h.name)
}

func main() {
	bruno := new(Human)
	bruno.name = "bruno kofi edoh"
	var myAge uint16 = 200
	fmt.Println(myAge)
	fmt.Println(bruno)
	fmt.Println(bruno.info())
}