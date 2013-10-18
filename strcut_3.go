package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int

	contact struct {
		phone    string
		address  string
		postcode string
	}
}

func test() {
	d := map[string]struct {
		x string
		y int
	}{
		"a": {"a10", 10},
		"b": {"b20", 20},
	}
	fmt.Println(d)
}

func main() {
	var d = struct {
		name  string
		age   int
		title string
	}{"user1", 10, "cto"}

	f := map[string]struct {
		x string
		y int
	}{
		"a": {"a10", 10},
		"b": {"b20", 20},
	}

	fmt.Println(f)

	d2 := Person{
		name: "user1",
		age:  10,
	}
	// t := &test()
	d2.contact.phone = "1350000000"
	d2.contact.address = "beijing"
	d2.contact.postcode = "100000"
	fmt.Println(d, "\n", d2)
	// fmt.Println(t)
}
