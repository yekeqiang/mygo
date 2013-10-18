package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	tick := time.Tick(1 * time.Second)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-tick:
				fmt.Println("tick")
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()

	<-o
}
