package main

import (
	"fmt"
)

func main() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	for i := 10; i > 0; i-- {
		if i == 5 {
			break //或者continue
		}
		fmt.Println(i)
	}
	fmt.Println("sum is equal to", sum)
	j := 10
	switch j {
	case 1:
		fmt.Println("j is equal to 1")
	case 2, 3, 4:
		fmt.Println("j is equal to 2,3 or 4")
	case 10:
		fmt.Println("j is equal to 10")
	default:
		fmt.Println("All I know is that j is an integer")
	}
}
