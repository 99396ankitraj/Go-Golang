package main

import (
	"fmt"
)

func main(){
	fmt.Println("Now i am trying to learn Map")
	mp := make(map[string]int)

	mp["anshika"] = 100
	mp["sontee"] = 98
	mp["asmat"] = 102
	mp["riya"] = 96

	fmt.Println("marks of anshika: ",mp["anshika"])
	mp["riya"]=99
	fmt.Println("New marks of Riya ",mp["riya"])

	for ind , val := range mp{
		fmt.Printf("key is %s and marks is %d\n",ind,val)
	}

	//cheacking if a key exists 
	Grades, Exists := mp["anshika"]
	fmt.Println("Grades of Prince: ",Grades)
	fmt.Println("anshika exists: ",Exists)
}