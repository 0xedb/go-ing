package main

import (
	"fmt"
)

// Ruler draws ruler on screen
func Ruler(til int) {
	fmt.Println("----")
	if til == 0 {
		return
	}

	fmt.Println("---")
	fmt.Println("--")
	fmt.Println("-")
	Ruler(til - 1)

}

func main() {
	Ruler(5)
}
