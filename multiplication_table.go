package main

import (
	"fmt"
)

func addStudent(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

func findStudent(grades map[string]int, name string) {
	grade, ok := grades[name]
	if ok {
		fmt.Printf("Оценка студента %s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func deleteStudent(grades map[string]int, name string) {
	delete(grades, name)
	fmt.Printf("Студент %s удалён\n", name)
}

func main() {
	grades := make(map[string]int)

	addStudent(grades, "Алиса", 5)
	addStudent(grades, "Боб", 4)

	findStudent(grades, "Алиса")
	findStudent(grades, "Ева")

	deleteStudent(grades, "Боб")
	findStudent(grades, "Боб")
}
