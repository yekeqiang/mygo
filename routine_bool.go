package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int, 3) //修改2为1就报错，修改2为3就正常运行
	c <- 1
	// runtime.GOMAXPROCS(2)
	c <- 2
	fmt.Println(<-c)

	fmt.Println(<-c)
}
