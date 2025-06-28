package main

import "fmt"

func isLeap(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if isLeap(year) {
		fmt.Println("Год високосный")
	} else {
		fmt.Println("Год обычный")
	}
}
