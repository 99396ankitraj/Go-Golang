package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("I am learning data conversion in GO lang")
	var num int = 42
	fmt.Println("Number is ", num)
	fmt.Println("Type of Data is %T\n", num)

	var data float64 = float64(num)
	fmt.Println("Data is :", data)
	fmt.Println("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "1234"
	number_int , _:= strconv.Atoi(number_string)
	fmt.Println("number_int is",number_int)
	fmt.Printf("Type of number_int is %T\n" , number_int)
}
