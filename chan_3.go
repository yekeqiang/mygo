package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("recv:", <-c)
		o <- true
	}()

	c <- 10000
	fmt.Println("send over...")
	<-o
}
