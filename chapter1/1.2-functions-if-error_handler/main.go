package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hi")

	print()
	printMessage("Message Hi")
	printMessage("Message Hi hi")

	gr := sayHello("User", 32)
	fmt.Println(gr)

	fmt.Println(enterTheClub(12))
	mes, entered := enterTheClub(21)
	fmt.Println(mes, entered)

	m, err := enterTheClub(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

func print() {
	fmt.Println("Function Hi")
}

func printMessage(m string) {
	fmt.Println(m)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s ! You are %d years old", name, age)
}

func enterTheClub(age int) (string, error) {
	if age < 18 {
		return "No !", errors.New("too young")
	}

	return "Yes", nil
}
