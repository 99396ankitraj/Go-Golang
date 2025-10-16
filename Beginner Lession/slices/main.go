package main

import (
	"fmt"
)

func main(){
	fmt.Println("Now we are going to learn about basics of slices")

	// unlike arrays slices have a dynamic size and theire length can be changed during the programs execution
	// basically slices are dynamic arrays

	// numbers := []int{1,2,3,4,5}
	// numbers = append(numbers , 3,10,12,13,14,15,16,17,18,19)
	// fmt.Println("Number : ",numbers)
	// fmt.Printf("Number has data type: %T\n",numbers)
	// fmt.Println("Length: ",len(numbers))


	//slice decelaration using slice 
	numbers := make([]int , 3 , 5)

	numbers = append(numbers , 4)
	numbers = append(numbers , 98)
	numbers = append(numbers , 98)

	fmt.Println("Slice:",numbers)
	fmt.Println("Length:",len(numbers))
	fmt.Println("capavity:", cap(numbers))


}