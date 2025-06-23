package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Go is a statically typed programming language"

	words := strings.Fields(sentence)
	for _, word := range words {
		fmt.Println(word)
	}
}
