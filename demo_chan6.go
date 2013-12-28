package main

import (
	"fmt"
)

func send(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1)
	go send(ch2)
	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("%s: %d\n", "CH1", v1)
		case v2 := <-ch2:
			fmt.Printf("%s: %d\n", "CH2", v2)
		}
	}
	fmt.Println("End.")
}
