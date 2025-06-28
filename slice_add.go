package main

import "fmt"

func main() {
	var fruits []string

	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "cherry")

	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}
}
