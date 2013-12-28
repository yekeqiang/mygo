package main

import (
	"fmt"
	// "time"
)

func main() {
	tick := make(chan int)
	go func() {
		// time.Sleep(2 * time.Second)
		count := 1
		for {
			// time.Sleep(3 * time.Second)
			tick <- count
			count++
		}
	}()

	for v := range tick {
		fmt.Printf("%d\n", v)
		if v == 5 {
			break
		}
	}
}
