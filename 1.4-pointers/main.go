package main

import (
	"errors"
	"fmt"
	"reflect"
)

var startMessage string

// Run always in package inicialisation
func init() {
	startMessage = "Hi. Pointers"
}

func main() {
	fmt.Println(startMessage)

	message := "Message !"
	fmt.Println(message)
	changeMessage(&message)
	fmt.Println(message)
	fmt.Println("--------------")

	var p *int
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)
	var number = 45
	p = &number
	fmt.Println(p)
	fmt.Println(&number)
	fmt.Println(number)
	*p = 333
	fmt.Println(number)
	fmt.Println("--------------")

	// ARRAY
	messagesAr := [3]string{"11", "22", "33"}
	fmt.Println(reflect.TypeOf(messagesAr))
	fmt.Println(messagesAr)
	fmt.Println(messagesAr[2])
	messagesAr[2] = "three"
	fmt.Println(messagesAr[2])
	printArray(messagesAr)
	fmt.Println("--------------")

	// SLICE
	messages := []string{"1", "2", "3"}
	fmt.Println(messages)
	printSlice(messages)
	fmt.Println(messages)
	var messeges2 []string
	fmt.Println(reflect.TypeOf(messeges2))
	fmt.Println(messeges2)
	messeges3 := make([]string, 5, 8)
	fmt.Println(reflect.TypeOf(messeges3))
	fmt.Println(messeges3)
	messeges3[1] = "2"
	fmt.Println(messeges3)
	fmt.Println(len(messeges3))
	fmt.Println(cap(messeges3))
	messeges3 = append(messeges3, "6")
	messeges3 = append(messeges3, "7")
	messeges3 = append(messeges3, "8")
	fmt.Println(messeges3)
	fmt.Println(len(messeges3))
	fmt.Println(cap(messeges3))
	messeges3 = append(messeges3, "9")
	fmt.Println(messeges3)
	fmt.Println(len(messeges3))
	fmt.Println(cap(messeges3))
	fmt.Println("--------------")

	matrix := make([][]int, 10)
	fmt.Println(matrix)
	counter := 0
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			counter++
			matrix[i][j] = counter
		}
		fmt.Println(matrix[i])
	}
	fmt.Println(matrix)
	fmt.Println("--------------")

	messages5 := []string{
		"Message 1",
		"Message 2",
		"Message 3",
		"Message 4",
		"Message 5",
		"Message 6",
	}
	fmt.Println(messages5)
	for i, val := range messages5 {
		fmt.Println(i)
		fmt.Println(val)
	}
	fmt.Println("--------------")

	counter2 := 0
	for {
		counter2++
		fmt.Println(counter2)
		if counter2 > 9 {
			break
		}
	}
}

func changeMessage(msg *string) {
	// fmt.Println(reflect.TypeOf(msg))
	// fmt.Println(msg)
	*msg += " From function printMessage()"
}

func printArray(messages [3]string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	fmt.Println(messages)

	return nil
}

func printSlice(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty slice")
	}

	messages[0] = "one"

	fmt.Println(messages)

	return nil
}
