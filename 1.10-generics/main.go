package main

import (
	"fmt"
)

var startMessage string

func init() {
	startMessage = "Hi. Generics"
}

func main() {
	fmt.Println(startMessage)
}
