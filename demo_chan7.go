package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()
	timeout := time.After(5 * time.Second)
	overTag := make(chan bool)
	go func() {
		for {
			select {
			case v1, ok := <-ch1:
				if !ok {
					overTag <- true
					break
				}
				fmt.Printf("%s: %d\n", "CH1", v1)
			case <-timeout:
				fmt.Println("Timeour.")
				overTag <- true
			}
		}
	}()
	<-overTag
	fmt.Println("End.")
}
