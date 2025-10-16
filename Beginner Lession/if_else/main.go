package main

import(
	"fmt"
)

func main(){
	fmt.Println("I am trying to learn If else statement")

	age := 49

	if age > 50{
		fmt.Println("User age is greater then 50")
	}else if age < 50{
		fmt.Println("User age is less then 50")
	}else{
		fmt.Println("User age is equal to 50")
	}
}