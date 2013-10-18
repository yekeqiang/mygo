package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func main() {
	u := &User{100, "Tom"}
	fmt.Println(u.Id, u.Name)
	u.Id = 200
	fmt.Println(u.Id, u.Name)

	m := User{300, "Jack"}
	fmt.Println(m.Id, m.Name)

	// fmt.Println(User{"100"} == User{300})
}
