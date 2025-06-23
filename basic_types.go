package main

import "fmt"

func main() {
	var (
		a int       = 42
		b float64   = 3.14
		c string    = "GoLang"
		d bool      = true
		e byte      = 'A'
		f rune      = 'Ж'
		g complex64 = 2 + 3i
	)

	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)
	fmt.Println("byte:", e)
	fmt.Println("rune:", f)
	fmt.Println("complex64:", g)
}
