package main

import (
	"fmt"
)

type Node struct {
	_     int
	Value string
	Next  *Node
}

func main() {
	a := &Node{Value: "a"}
	b := &Node{Value: "b", Next: a}

	fmt.Printf("%#v\n", a)
	fmt.Println("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Println("%#v\n", b)
}
