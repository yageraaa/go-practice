package main

import (
	"fmt"
	"math"
)

// Интерфейс
type Shape interface {
	Area() float64
}

// Прямоугольник
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Круг
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
}

func main() {
	r := Rectangle{Width: 4, Height: 5}
	c := Circle{Radius: 3}

	printArea(r)
	printArea(c)
}
