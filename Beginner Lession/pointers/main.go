package main

import (
	"fmt"
)

func main(){
	fmt.Println("now i am going to learn about the pointers concept in the golang")
	var num int
	num = 2
	var ptr *int
	ptr = &num
	fmt.Println("Num has value: ",num)
	fmt.Println("Pointers Contain: ",*ptr)



	//in sort we can do this as 
	num1 := 2 
	ptr1 := &num
	fmt.Println("Num has value: ",num1)
	fmt.Println("Pointers Contain: ",*ptr1)

}
