package main

import (
	"fmt"
)

func test() (int, string) {
	return 123, "abc"
}
func main() {
	a, _ := test()
	fmt.Println(a)
}
