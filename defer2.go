package main

import (
	"fmt"
)

func doubleSum(a, b int) (sum int) {
	defer func() {
		sum *= 2
	}()
	sum = a + b
	return sum
}

func main() {
	// doubleSum(2, 3)
	fmt.Println(doubleSum(2, 3))
}
