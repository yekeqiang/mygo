package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello ", name, ", my name is ", u.name)
}

func main() {
	u := User{1, "ok", 12}
	v := reflect.ValueOf(u)
	// mv := v.MethdByName("Hello")
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("tom")}
	mv.Call(args)
}
