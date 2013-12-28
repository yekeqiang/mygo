package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2) //如果这里不设置容量的话，那下面的channel就会堵塞。导致执行default
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	//messages <- "hi"
	//fmt.Println(msg)
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	//messages <- "hello"
	//messages := <-"hello"
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
