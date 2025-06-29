package main

import "fmt"

type Student struct {
    Name       string
    Age        int
    Course     int
    AvgGrade   float64
}

func printStudent(s Student) {
    fmt.Printf("Имя: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n", s.Name, s.Age, s.Course, s.AvgGrade)
}

func updateGrade(s *Student, newGrade float64) {
    s.AvgGrade = newGrade
}

func main() {
    student := Student{"Иван", 20, 3, 4.2}
    printStudent(student)

    updateGrade(&student, 4.8)
    fmt.Println("\nПосле обновления оценки:")
    printStudent(student)
}
