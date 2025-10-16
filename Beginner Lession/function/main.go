package main

import (
	"fmt"
)

func solve(){
	fmt.Println("HII my name is ankit raj")
}

func add(a , b , c int) int {
	return a + b + c
}

func main(){
	fmt.Println("Hii , we are learning functions in the go")
	solve()
	sum := add(4 , 5 , 6)
	fmt.Print("addition of two number is ", sum)
}