package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World! Welcome to  to to to Go programming."
	words := strings.Split(str , " ")
	fmt.Println(words)
	count := strings.Count(str, "to")
	fmt.Println("Count of 'to':", count)
	trimmed := strings.TrimSpace("   Hello, Go!   ")
	fmt.Println("Trimmed string:", trimmed)
}