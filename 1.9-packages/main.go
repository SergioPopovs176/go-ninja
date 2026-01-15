package main

import (
	"fmt"
	"go-ninja/basic/shape"
	"time"

	"github.com/zhashkevych/scheduler"
)

var startMessage string

func init() {
	startMessage = "Hi. Packages"
}

func main() {
	fmt.Println(startMessage)

	s := shape.NewSquare(6)
	printShapeAreaAndPerimeter(s)

	c := shape.NewCircle(5)
	printShapeAreaAndPerimeter(c)

	sch := scheduler.NewScheduler()
	fmt.Println(sch)

	t := time.Now()
	fmt.Println(t)
}

func printShapeAreaAndPerimeter(shape shape.Shape) {
	fmt.Println("Area", shape.Area(), "Perimeter", shape.Perimetr())
}
