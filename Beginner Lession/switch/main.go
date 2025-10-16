package main

import (
	"fmt"
)

func main(){
	fmt.Println("I am trying to learn Switch statement")

	temperature := 10

	switch{
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature <10 :
		fmt.Println("cold")
	case temperature >= 20 && temperature < 30:
		fmt.Println("warm")
	default:
		fmt.Println("hot")
	}
}