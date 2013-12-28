package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				select {
				case ch1 <- 1:
					// case ch2 <- 2:
					// case ch3 <- 3:
				}
			case <-time.After(2 * time.Second):
				fmt.Println("Time out.Stooped ticker.")
				close(ch1)
				return
			}
		}
	}()
	overTag := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					fmt.Println("Closed channel")
					overTag <- true
					break
				}
				fmt.Printf("%s: %d\n", "CH1", v)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Stop ticker")
	ticker.Stop()
	<-overTag
	fmt.Println("End.")

}
