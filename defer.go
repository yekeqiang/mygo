package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Return normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", f)
		}
	}()
	fmt.Println("Calling g.")
	g()
	fmt.Println("Returned normally from g.")

}

func g() {
	panic("ERROR")
}
