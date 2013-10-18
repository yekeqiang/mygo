package main

import (
	"fmt"
	_ "runtime"
	str "strings"
	"time"
)

func init() {
	fmt.Println(a)
	fmt.Println("init Over...")
}

var a string = str.Join([]string{"a", "b"}, ",")

func init() {
	time.Sleep(2 * time.Second)
	fmt.Println("init2 over...")
}

func main() {
	fmt.Println("Hello,World!")
}
