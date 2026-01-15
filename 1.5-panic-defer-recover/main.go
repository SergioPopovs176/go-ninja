package main

import (
	"fmt"
)

var startMessage string

// Run always in package inicialisation
func init() {
	startMessage = "Hi. Panic"

}

func main() {
	defer handlerPanic()

	fmt.Println(startMessage)

	fmt.Println("main()")

	messages := []string{
		"Message 1",
		"Message 2",
		"Message 3",
		"Message 4",
	}
	messages[4] = "message"
	fmt.Println(messages)
}

func handlerPanic() {
	fmt.Println("Recover panic !")
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
