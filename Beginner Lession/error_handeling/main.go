package main

import(
	"fmt"
)

func divide(a , b float64) (float64 , error){
	if (b == 0){
		return 0 , fmt.Errorf("denminator must not be zero")
	}

	return a/b , nil
}
func main(){
	fmt.Println("i am trying to impeament error handelling")

	ans , error := divide(10 , 0)

	if(error != nil){
		fmt.Println(error)
	}

	fmt.Println("Answer of divide is",ans);

}