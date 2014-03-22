package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int
	Name string `sql:"unique"`
	Age  int
}

func main() {
	db, _ := gorm.Open("sqlite3", "test.db")
	db.LogMode(true)

	db.CreateTable(new(User))

	userOne := User{Name: "Ann", Age: 101}
	userTwo := User{Name: "Bob", Age: 102}

	db.Save(&userOne)
	db.Save(&userTwo)

	var bob User
	db.Where("name = ?", "Bob").Find(&bob)
}
