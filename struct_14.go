package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

type Manager struct {
	Group string
	User
}

type IUser interface {
	Test()
}

type IManager interface {
	Test()
	Test2()
}

func (this *User) Test() {
	fmt.Println(this)
}

func (this *User) Test2() {
	fmt.Println(this)
}

func main() {
	var im IManager = &Manager{"IT", User{1, "Tom"}}
	im.Test()
	im.Test2()

	var iu IUser = im
	iu.Test()
}
