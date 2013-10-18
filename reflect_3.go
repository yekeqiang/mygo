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

func (this User) Test(x int) string {
	return fmt.Sprintf("user.name:%s;x:%d", this.Name, x)
}

func main() {
	var o interface{} = &User{1, "Tom", 30, "Programmer"}
	v := reflect.ValueOf(o)
	m := v.MethodByName("Test")
	// m2 := reflect.ValueOf(1234)
	m3 := []reflect.Value{reflect.ValueOf(1234)}
	fmt.Println(m.Call(m3))
	fmt.Println(m3)

	rets := m.Call([]reflect.Value{reflect.ValueOf(1234)})
	fmt.Println(rets)
	s := rets[0].Interface().(string)
	fmt.Printf("%T,%v\n", s, s)
}
