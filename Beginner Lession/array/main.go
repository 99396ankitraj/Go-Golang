package main

import (
	"fmt"
)

func main(){
	fmt.Println("now i am going to work with the array")

	var name[5] string
	name[0] = "anshika"
	name[1] = "sontee"
	name[2] = "asmat"
	name[3] = "riya"
	name[4] = "sweety"
	fmt.Println("The complete array is ",name)


	var numbers = [5]int{1,2,3,4,5}
	fmt.Println("The complete array is ",numbers)

	fmt.Println("Length of Numbers array is :", len(numbers))

	var name1[5] int
	fmt.Println("lets print the empty array",name1 )

	var name2[5] string
	fmt.Println("lets print the empty array",name2)
	
	var price[5] string
	fmt.Printf("Price array is %q\n",price)

}