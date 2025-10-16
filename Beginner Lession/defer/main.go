package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World")
	defer fmt.Println("from")
	fmt.Println("Hello")
}


//it follows LIFO (Last In First Out) order
//So if there are multiple defer statements, they will be executed in reverse order of their appearance in the code. 