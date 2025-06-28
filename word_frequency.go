package main

import "fmt"

func removeIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	items := []string{"a", "b", "c", "d", "e"}
	fmt.Println("До удаления:", items)

	indexToRemove := 2
	items = removeIndex(items, indexToRemove)

	fmt.Println("После удаления:", items)
}
