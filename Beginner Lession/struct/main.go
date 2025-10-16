package main

import (
	"fmt"
)

type Person struct{
	firstName string
	lastName string 
	age int
}

type Contact struct{
	Email string
	Phone string
}

type Address struct{
	house int
	area string
	state string
}

type Employee struct{
	person_details Person
	person_contact Contact
	person_address Address
}

func main(){
	fmt.Println("Now i am trying to learn Struct")

	//first method
	// var ankit Person
	// ankit.firstName = "anshika"
	// ankit.lastName = "bhardwaj"
	// ankit.age = 39

	// second method
	// Person1 := Person{
	// 	firstName: "sontee",
	// 	lastName: "bhardwaj",
	// 	age: 22,
	// }

	//with the help of new keyword
	// var Person2 = new(Person)
	// Person2.firstName = "riya"
	// Person2.lastName = "kashyap"
	// Person2.age = 23

	// fmt.Println("Ankit Person : ", Person2)
	// fmt.Println("Age of Persion 2 is : ", Person2.age)



	//now we will work on complex struct

	 var Employee1 Employee
	 Employee1.person_details = Person{
		firstName: "ankit",
		lastName: "raj",
		age: 24,
	 }
	 Employee1.person_contact.Email="ankitraj@gmail.com"
	 Employee1.person_contact.Phone="9955339365"

	 Employee1.person_address = Address{
		house: 12345,
		area:"aurangbad bihar",
		state:"bihar",

	 }
	 fmt.Println("Employee 1 : ",Employee1)
}