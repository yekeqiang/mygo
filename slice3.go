package main

import (
	"fmt"
)

func modify(a *int) {
	a = nil
}

func main() {
	a := new(int)
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}
