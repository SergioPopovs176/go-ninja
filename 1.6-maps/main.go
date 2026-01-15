package main

import (
	"fmt"
)

var startMessage string

// Run always in package inicialisation
func init() {
	startMessage = "Hi. Maps"
}

func main() {
	fmt.Println(startMessage)

	users := map[string]int{
		"Vasja": 12,
		"Kosta": 15,
		"Ivan":  9,
	}
	fmt.Println(len(users))
	fmt.Println("--------------")

	fmt.Println(users)
	fmt.Println(users["Ivan"])
	age, exist := users["Vasja"]
	fmt.Println(age)
	fmt.Println(exist)
	age, exist = users["Sergio"]
	fmt.Println(age)
	fmt.Println(exist)
	fmt.Println("--------------")

	for name, age := range users {
		fmt.Println(name, "--", age)
	}
	fmt.Println("--------------")

	users["Sergio"] = 44
	for name, age := range users {
		fmt.Println(name, "--", age)
	}
	fmt.Println("--------------")

	delete(users, "Ivan")
	for name, age := range users {
		fmt.Println(name, "--", age)
	}
	fmt.Println("--------------")

	var users2 map[string]int
	fmt.Println(users2)
	fmt.Println("--------------")

	users3 := make(map[string]int)
	fmt.Println(users3)
	users3["Ivanchik"] = 56
	fmt.Println(users3)
	fmt.Println("--------------")
}
