package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Title string
}

func (this User) Test(x int) {
	fmt.Println("User.Test", this.Name, x)
}

func test(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println(k)
		return
	}

	fmt.Println("Type:", t.Name())
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf(" %6s: %v = %v\n", field.Name, field.Type, value)
	}

	fmt.Println("Method:")

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf(" %s:%v\n", method.Name, method.Type)
	}
}

func main() {
	u := User{1, "Tom", 30, "Programmer"}
	test(u)
}
