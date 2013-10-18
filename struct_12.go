package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

type Tester interface {
	Test()
}

func (this *User) Test() {
	fmt.Println("*User.Test,", this)
}

func doSomething(i interface{}) {
	if o, ok := i.(Tester); ok {
		o.Test()
	}

	i.(*User).Test()
	// i.User.Test()
}

func main() {
	u := &User{1, "Tom"}
	doSomething(u)
}
