package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
		close(ch1)
	}()
	tick := time.Tick(1 * time.Second)
	overTag := make(chan bool)
	go func() {
		for {
			select {
			case <-tick:
				v, ok := <-ch1
				if !ok {
					overTag <- true
					fmt.Println("closed channel.")
					break
				}
				fmt.Printf("%s: %d\n", "CH1", v)
			}
		}
	}()
	<-overTag
	fmt.Println("End.")
}
