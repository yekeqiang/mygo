package main

import (
	"fmt"
)

func modifySliceData(data []int) {
	data[1] = 3
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	modifySliceData(a)
	fmt.Println(a)
}
