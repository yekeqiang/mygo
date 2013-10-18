package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("start:", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n, runtime.NumGoroutine())
		}(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("over:", runtime.NumGoroutine())
}
