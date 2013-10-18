package main

import (
	"fmt"
)

type Myint int
type InPointer *int

type Data struct {
	x, y int
}

type ITest interface {
	test(int) (string, int)
}

func main() {
	var x = 123
	var p InPointer = &x
	fmt.Println(*p)

	var d = Data{1, 2}
	fmt.Println(d.x, d.y)

	p2 := &Data{y: 200}
	fmt.Println(p2.x, p2.y)
}
