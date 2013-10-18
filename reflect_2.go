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

func test(i interface{}) {
	v := reflect.ValueOf(i)

	//首先判断是Ptr类型
	//用Elem()获取ptr指向的实际对象，并判断是settable
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("Cannot Set!")
		return
	} else {
		v = v.Elem() //我们要操作的对象，而不是ptr
	}

	field := v.FieldByName("Name")
	if field.Kind() == reflect.String {
		field.SetString("Jack")
	}
}

func main() {
	u := User{1, "Tom", 30, "Programmer"}
	test(&u)
	// test(u)
	fmt.Println(u)
}
