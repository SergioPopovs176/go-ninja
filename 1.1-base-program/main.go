package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "Hi, all !!!"
	fmt.Println(message)
	message = "Hi, all !!! again"
	fmt.Println(message)

	var otherMessage string
	otherMessage = "Other message"
	fmt.Println(otherMessage)

	const MESSAGE string = "Const message"
	fmt.Println(MESSAGE)

	fmt.Println(reflect.TypeOf(message))

	// Defalt value for empty variable
	var mes string
	var number int
	var b bool
	var r rune
	fmt.Println(mes)
	fmt.Println(number)
	fmt.Println(b)
	fmt.Println(r)

	r = 'g'
	fmt.Println(r)

	a1, a2, a3 := 11, 22, 33
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a1, a2, a3)
	a1, a2 = a2, a1
	fmt.Println(a1, a2, a3)
	a1, _, a3 = 1, 77, 3
	fmt.Println(a1, a2, a3)
}
