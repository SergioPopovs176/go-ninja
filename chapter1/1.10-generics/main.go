package main

import (
	"fmt"
)

var startMessage string

func init() {
	startMessage = "Hi. Generics"
}

type Number interface {
	int64 | float32
}

type User struct {
	name  string
	email string
}

func main() {
	fmt.Println(startMessage)

	ints := []int64{1, 2, 3, 4, 5}
	floats := []float32{1.4, 6.3, 0.45}

	fmt.Println(ints)
	fmt.Println(sumInts(ints))
	fmt.Println(sum(ints))
	fmt.Println(floats)
	fmt.Println(sumFloats(floats))
	fmt.Println(sum(floats))

	r := searchElement(ints, 3)
	fmt.Println(r)
	r = searchElement(ints, 34)
	fmt.Println(r)

	strs := []string{"odin", "blue", "oK"}
	fmt.Println(strs)
	r = searchElement(strs, "OK")
	fmt.Println(r)

	users := []User{
		{name: "Ivan", email: "ivan@gmail.com"},
		{name: "Vasja", email: "vas@gmail.com"},
		{name: "Givi", email: "givi@gmail.com"},
	}
	fmt.Println(users)
	r = searchElement(users, User{name: "Ivans", email: "ivan@gmail.com"})
	fmt.Println(r)

	printAny(ints)
	printAny(users)
}

// func sum[V int64 | float32](input []V) V {
// 	var sum V
// 	sum = 0
// 	for _, v := range input {
// 		sum += v
// 	}

//		return sum
//	}
//
// OR
func sum[V Number](input []V) V {
	var sum V
	sum = 0
	for _, v := range input {
		sum += v
	}

	return sum
}

func sumInts(input []int64) int64 {
	var sum int64
	sum = 0
	for _, v := range input {
		sum += v
	}

	return sum
}

func sumFloats(input []float32) float32 {
	var sum float32
	sum = 0
	for _, v := range input {
		sum += v
	}

	return sum
}

func searchElement[C comparable](elements []C, element C) bool {
	for _, v := range elements {
		if v == element {
			return true
		}
	}

	return false
}

func printAny[A any](input A) {
	fmt.Println(input)
}
