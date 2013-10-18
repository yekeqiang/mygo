package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("go .....")
	}()

	fmt.Println("hi!")
	time.Sleep(2 * time.Second)
}
