package main

import (
	"fmt"
	// "runtime"
)

func test(c chan bool, b bool) {
	x := 0
	for i := 0; i < 100000000; i++ {
		x += i
	}

	fmt.Println(x)
	if b {
		c <- true
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	c := make(chan bool)

	for i := 0; i < 100; i++ {
		go test(c, i == 99)
	}

	<-c
}
