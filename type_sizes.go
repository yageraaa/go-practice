package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("string: %d bytes\n", unsafe.Sizeof(""))
	fmt.Printf("bool: %d bytes\n", unsafe.Sizeof(true))
	fmt.Printf("byte: %d bytes\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("rune: %d bytes\n", unsafe.Sizeof(rune(0)))
	fmt.Printf("complex64: %d bytes\n", unsafe.Sizeof(complex64(0)))
}
