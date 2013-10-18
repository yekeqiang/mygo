package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id   int
	Name string
}

type Manager struct {
	Group string
	User
}

func main() {
	m := Manager{User: User{1, "Tom"}, Group: "IT"}

	fmt.Printf("m alignof = %d\n", unsafe.Alignof(m))
	fmt.Printf("m address = %p\n", &m)
	fmt.Printf("m size = %d\n", unsafe.Sizeof(m))

	fmt.Printf("m.Group address = %p\n", &m.Group)
	fmt.Printf("m.Group offset = %d\n", unsafe.Offsetof(m.Group))
	fmt.Printf("m.Group size =%d\n", unsafe.Sizeof(m.Group))

	fmt.Printf("m.User address = %p\n", &m.Group)
	fmt.Printf("m.User offset = %d\n", unsafe.Offsetof(m.User))
	fmt.Printf("m.User size = %d\n", unsafe.Sizeof(m.User))

	fmt.Printf("User.Id address = %p\n", &m.User.Id)
	fmt.Printf("User.Id offset = %d\n", &m.User.Id)
	fmt.Printf("User.Id size = %d\n", &m.User.Id)

	fmt.Printf("User.Name address = %p\n", &m.User.Name)
	fmt.Printf("User.Name offset = %d\n", &m.User.Name)
	fmt.Printf("User.Name size = %d\n", &m.User.Name)

}
