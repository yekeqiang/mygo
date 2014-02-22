package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, _ := os.Create("ykq.txt")
	defer f.Close()
	fmt.Println(reflect.ValueOf(f).Type())
}
