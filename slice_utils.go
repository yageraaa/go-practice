package main

import (
	"fmt"
	"sort"
)

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func filterEven(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{5, 2, 8, 3, 9, 1, 4}

	fmt.Println("Содержит 3?", contains(nums, 3))

	sort.Ints(nums)
	fmt.Println("Отсортированный:", nums)

	evens := filterEven(nums)
	fmt.Println("Чётные:", evens)
}
