package main

import (
	"fmt"
	"time"
)

func task(args ...int) chan int {
	x := 0
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)

		for _, v := range args {
			x += v
		}

		c <- x
	}()
	return c
}

func main() {
	c := task(1, 2, 3, 4)
	fmt.Println("do something")
	i := <-c
	fmt.Println("task result = ", i)
}
