package main

var a string
var c = make(chan int, 10)

func f() {
	a = "hello world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
