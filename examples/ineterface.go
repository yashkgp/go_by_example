package main

import (
	"fmt"
	Math "math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return Math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * Math.Pi * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (r Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func get_shape(s Shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	c := Circle{x: 0, y: 0, radius: 5}
	r := Rectangle{width: 10, height: 5}
	fmt.Println("Circle area: ", c.area(), "Circle perimeter: ", c.perimeter())
	fmt.Println("Rectangle area: ", r.area(), "Rectangle perimeter: ", r.perimeter())
	shapes := []Shape{c, r}
	for _, shape := range shapes {
		fmt.Println(shape.area())
		fmt.Println(shape.perimeter())
	}
	get_shape(c)
}
