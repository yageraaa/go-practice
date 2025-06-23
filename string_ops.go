package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, GoLang World!"

	fmt.Println("Длина строки:", len(str))

	substr := "GoLang"
	found := strings.Contains(str, substr)
	fmt.Printf("Подстрока \"%s\" найдена: %v\n", substr, found)
	fmt.Println("Верхний регистр:", strings.ToUpper(str))
	fmt.Println("Нижний регистр:", strings.ToLower(str))
}
