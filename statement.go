package main

import (
	"fmt"
)

func main() {
	i := 0
	p := &i
	//a := i++ //syntax error:unexpected ++,
	// i++
	i++
	a := i

	*p++
	b := *p

	fmt.Println(a, b)
}
