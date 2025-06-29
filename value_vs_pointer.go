package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func changeByValue(p Person) {
    p.Age = 99
}

func changeByPointer(p *Person) {
    p.Age = 99
}

func main() {
    p := Person{Name: "Алиса", Age: 25}

    changeByValue(p)
    fmt.Println("После changeByValue:", p.Age) // останется 25

    changeByPointer(&p)
    fmt.Println("После changeByPointer:", p.Age) // станет 99
}
