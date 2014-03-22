package main

import (
	"fmt"
)

func main() {
	data := make(chan int, 4)
	exit := make(chan bool)

	data <- 1
	data <- 2
	data <- 3
	data <- 4

	go func() {
		for d := range data {
			fmt.Println(d)
		}

		exit <- true
	}()

	data <- 4
	data <- 5
	close(data)

	<-exit
}
