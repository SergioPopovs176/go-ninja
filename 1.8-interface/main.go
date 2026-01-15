package main

import (
	"fmt"
	"math"
)

var startMessage string

func init() {
	startMessage = "Hi. Interface"
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimetr() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

func (s Square) Perimetr() float32 {
	return s.sideLength * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimetr() float32 {
	return 2 * c.radius * math.Pi
}

func main() {
	fmt.Println(startMessage)

	s := Square{6}
	printShapeAreaAndPerimeter(s)

	c := Circle{5}
	printShapeAreaAndPerimeter(c)

	printInterface(c)
	printInterface(s)
	printInterface(44)
	printInterface("dfhhfhdsk")
}

func printShapeAreaAndPerimeter(shape Shape) {
	fmt.Println("Area", shape.Area(), "Perimeter", shape.Perimetr())
}

func printInterface(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("intiger", value)
	case bool:
		fmt.Println("boolian", value)
	default:
		fmt.Println("undefined type", value)
	}

	str, ok := i.(string)
	if ok {
		fmt.Println("String length", len(str))
	}

	// fmt.Printf("%+v\n", i)
}
