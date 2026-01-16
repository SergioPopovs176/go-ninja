package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hi. Functions 2")

	result := prediction("w")
	fmt.Println(result)

	min := findMin(5, 6, 78, -1, 7)
	fmt.Println(min)

	func() {
		fmt.Println("Anonim function")
	}()

	count := increment()
	fmt.Println(reflect.TypeOf(count))
	fmt.Println(count())
	fmt.Println(count())
}

func prediction(dayOfWeek string) string {
	switch dayOfWeek {
	case "mon":
		return "Good monday !"
	case "tu":
		return "Good tuesday !"
	case "w":
		return "Good wensday !"
	case "th":
		return "Good thuesday !"
	case "fr":
		return "Good friday !"
	case "sat":
		return "Good saturday !"
	case "sun":
		return "Good sunday !"
	}

	return ""
}

func findMin(numbers ...int) int {
	fmt.Println(reflect.TypeOf(numbers))
	fmt.Println(numbers)

	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, val := range numbers {
		if val < min {
			min = val
		}
	}

	return min
}

func increment() func() int {
	counter := 0

	return func() int {
		counter++
		return counter
	}
}
