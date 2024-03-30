// i want to calculate area, perimeter of geometry figure

package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length, Breadth float64
}

type Square struct {
	Length float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func NewRectangle(length, breadth float64) *Rectangle {
	return &Rectangle{length, breadth}
}

func NewSquare(length float64) *Square {
	return &Square{length}
}

func (a Circle) Area() float64 {
	return math.Pi * math.Pow(a.Radius, 2)
}

func (a Circle) Perimeter() float64 {
	return 2 * a.Radius
}

func (a Rectangle) Area() float64 {
	return a.Length * a.Breadth
}

func (a Rectangle) Perimeter() float64 {
	return 2 * (a.Length + a.Breadth)
}

func (a Square) Area() float64 {
	return math.Pow(a.Length, 2)
}

func (a Square) Perimeter() float64 {
	return 4 * a.Length
}

func GeometryFun(g []Geometry) {

	for _, shape := range g {
		fmt.Println(shape)
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}
}

func main() {

	circle := NewCircle(4.0)
	rectangle := NewRectangle(4, 16)
	square := NewSquare(30)
	g := []Geometry{circle, rectangle, square}
	fmt.Println("Calculate geometry of shape.")
	GeometryFun(g)
}