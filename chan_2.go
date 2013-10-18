package main

import (
	"fmt"
)

func recv(c <-chan int, over chan<- bool) {
	for v := range c {
		fmt.Println(v)
	}

	over <- true
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	o := make(chan bool)

	go send(c)
	go recv(c, o)

	<-o
}
