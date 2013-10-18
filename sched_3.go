package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("go1 defer...")

		for i := 0; i < 10; i++ {
			i += 1
			fmt.Println("go1: ", i)
			if i == 5 {
				runtime.Goexit()
			}
		}
	}()

	go func() {
		fmt.Println("go2")
	}()

	time.Sleep(3 * time.Second)
}
