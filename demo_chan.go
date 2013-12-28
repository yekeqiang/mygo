package main

import (
	"fmt"
)

var ch = make(chan int)
var content string

func set() {
	content = "It's a unbuffered chan!"
	<-ch
}

func main() {
	go set()
	ch <- 2
	fmt.Println(content)
}
