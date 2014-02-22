package main

import (
	"fmt"
	"time"
)

func main() {
	ok := make(chan bool, 1)
	go runLongProcess(ok)

	fmt.Println("Start waiting")
	<-ok
	fmt.Println("Data received")

}

func runLongProcess(ok chan bool) {
	time.Sleep(time.Second * 1)
	ok <- true
}
