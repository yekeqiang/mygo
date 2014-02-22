package main

import (
	"fmt"
)

func main() {
	var val interface{} = "good"
	fmt.Println(val.(string))
	// fmt.Println(val.(int))
}
