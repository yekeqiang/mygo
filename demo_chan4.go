package main

import (
	"fmt"
)

func receive(c <-chan int, over chan<- bool) {
	for v := range c {
		fmt.Println(v)
	}

	over <- true

}

func send(c chan<- int) {
	for i := 0; i < 3; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	o := make(chan bool)
	go receive(c, o)
	go send(c)
	<-o
}
