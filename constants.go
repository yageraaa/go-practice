package main

import (
	"fmt"
	"math"
)

const (
	Pi = math.Pi
	E  = math.E
)

func main() {
	radius := 3.0
	fmt.Printf("Площадь круга: %.2f\n", Pi*radius*radius)
	fmt.Printf("Значение числа e: %.4f\n", E)
}
