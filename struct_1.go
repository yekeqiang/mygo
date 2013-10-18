package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func main() {
	user1 := User{1, "Tom"}
	user2 := User{Name: "Jack"}

	fmt.Println(user1.Id, user1.Name)
	fmt.Println(user2.Id, user2.Name)
}
