package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func main() {
	u := User{100, "Tom"}
	fmt.Println(u.Id, u.Name)

	var u2 *User = new(User)
	*u2 = u
	fmt.Println(u2.Id, u2.Name)
}
