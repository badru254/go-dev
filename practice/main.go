package main

import (
	"fmt"
)

const aConst int = 645

func main() {
	//Explicit Typing
	var aString string = "This is a go var"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)

	var anInteger int = 45
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	//Implicit Typing
	var anotherString = "This is another gor var"
	fmt.Println(anotherString)
	fmt.Printf("The variable type is %T\n", anotherString)

	var anotherInt = 767
	fmt.Println(anotherInt)
	fmt.Printf("The variable type is %T\n", anotherInt)

	//Colon-equals operator
	myString := "this is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable type is %T\n", myString)

	//Constant
	fmt.Println(aConst)
	fmt.Printf("The variable type is %T\n", aConst)
}
