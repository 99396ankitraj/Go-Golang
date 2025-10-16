package main

import (
	"fmt"
)

func main(){
	fmt.Println("Now i am trying to learn for loops in GO")

	//normal for loops

	for i := 0 ; i < 3 ; i++{
		fmt.Println("Current index is :", i)
	}

	//if we want to run infinite loops in the go 

	counter := 0

	for{
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3{
			break
		}
	}

	//now we will learn range or iteration over string slices and arrays

	numbers := []int{11 , 43 , 34 , 54 , 65 , 65}
	for index, value := range numbers{
		fmt.Println("Index of Numbers is %d , and value is %d\n",index , value)
	}
}