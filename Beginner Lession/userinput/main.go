package main

import (
	"bufio"
	"fmt"
	"os"
)



func main(){
	// fmt.Println("Hey, What's your name ?")
	// var name string;
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name);


	// input including white spaces

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("hell0 Mr." , name)

}