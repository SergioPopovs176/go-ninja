package main

import (
	"fmt"
	"reflect"
)

var startMessage string

func init() {
	startMessage = "Hi. Structs"
}

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func newUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    Age(age),
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func (u User) printInfo() {
	fmt.Println(u.name, u.age, u.age.isAdult())
}

func (u *User) setName(name string) {
	u.name = name
}

type DumbDatabase struct {
	m map[string]string
}

func newDb() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func main() {
	fmt.Println(startMessage)

	user := User{"Vasja", 23, "Male", 76, 172}
	user2 := User{"Ivan", 44, "Male", 75, 150}
	fmt.Println(reflect.TypeOf(user))
	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", user2)
	user3 := newUser("Gosha", "Male", 12, 45, 110)
	fmt.Printf("%+v\n", user3)
	fmt.Println("--------------")

	db := newDb()
	fmt.Printf("%+v\n", db)
	fmt.Println("--------------")

	printUserInfo(user)
	user.printInfo()
	fmt.Println("--------------")

	fmt.Printf("%+v\n", user2)
	user2.setName("Toljan")
	fmt.Printf("%+v\n", user2)
	fmt.Println("--------------")

	var age Age = 56
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(age)
	fmt.Println(age.isAdult())
	fmt.Println("--------------")

}

func printUserInfo(user User) {
	fmt.Println(user.name, user.age)
}
