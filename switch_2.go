package main

import (
	"fmt"
)

func main() {
	// a := 5
	switch a := 5; {
	case a > 6:
		fmt.Println("a")
	case a > 9:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
