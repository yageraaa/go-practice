package main

import "fmt"

func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    x, y := 10, 20
    fmt.Printf("До обмена: x = %d, y = %d\n", x, y)

    swap(&x, &y)
    fmt.Printf("После обмена: x = %d, y = %d\n", x, y)
}
