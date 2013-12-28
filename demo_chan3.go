package main

func receiver(over chan<- bool) {
	over <- true
}

func main() {
	o := make(chan bool)
	go receiver(o)
	<-o
}
