package main

import "fmt"

func longestString(strs ...string) string {
    longest := ""
    for _, s := range strs {
        if len(s) > len(longest) {
            longest = s
        }
    }
    return longest
}

func main() {
    result := longestString("Go", "Gopher", "Programming", "Language")
    fmt.Println("Самая длинная строка:", result)
}
