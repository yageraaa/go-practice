package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Герман"
	currentDate := time.Now().Format("2006-01-02")
	fmt.Printf("Привет, %s! Сегодня: %s\n", name, currentDate)
}
