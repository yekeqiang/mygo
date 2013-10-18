package main

import (
	"fmt"
)

type MyInt int

type MyIntSlice []int

func main() {
	var a MyInt = 10

	var b int = int(a)

	var c MyInt = MyInt(b)

	fmt.Println(a, b, c)
}
