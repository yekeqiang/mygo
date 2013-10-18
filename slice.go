package main

import (
	"fmt"
)

func modifySlice(data []int) {
	data[0] = 3
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	modifySlice(a)
	fmt.Println(a)
}
