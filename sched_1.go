package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("go1:", i)
			if i == 5 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		fmt.Println("go2")
	}()

	time.Sleep(3 * time.Second)
}
