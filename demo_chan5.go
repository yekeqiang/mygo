package main

import (
	"fmt"
)

func send(ch chan<- int, number int) {
	ch <- number
	// close(ch)
}

func main() {
	// ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// go send(ch1, 5)
	go send(ch2, 2)
	go send(ch3, 3)

	select {
	case v3 := <-ch3:
		fmt.Printf("%s: %d\n", "CH1", v3)
	case v2 := <-ch2:
		fmt.Printf("%s: %d\n", "CH2", v2)
	case v3 := <-ch3:
		fmt.Printf("%s: %d\n", "CH3", v3)
		// default:
		// 	fmt.Printf("%s\n", "Hello")

	}
	fmt.Println("End.")
}
