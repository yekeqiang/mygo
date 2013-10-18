package main

import (
	"fmt"
)

type (
	Customer struct {
		Name   string
		Street []string
		City   string
		State  string
		Zip    string
	}
	Item struct {
		Id       int
		Name     string
		Quantity int
	}
	Items []Item
	Order struct {
		Id       int
		Customer Customer
		Items    Items
	}
)

func main() {
	// order = Order{Id:3;Customer:;Items:}
	fmt.Printf("%+v\n\n", order)
	fmt.Printf("%#v\n\n", order)
	fmt.Printf("%v\n\n", order)
	fmt.Printf("%s\n\n", order)
	fmt.Printf("%T\n", order)
}
