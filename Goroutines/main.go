package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, worls!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi prince")
	time.Sleep(250 * time.Millisecond)
	fmt.Println("Hi Ayan")

}

func main() {
	fmt.Println("Now i am learmning Goroutines..")
	go sayHello()
	go sayHi()
}
