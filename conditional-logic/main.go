package main

import (
	"fmt"
)

func main() {

	theAnswer := 0
	var resultMessage string

	//Condition not wrapped by parenthesis
	if theAnswer < 0 {
		resultMessage = "Less Than zero"
	} else if theAnswer > 0 {
		resultMessage = "More Than zero"
	} else {
		resultMessage = "Equal to Zero"
	}

	//You can initialize variable inside if condition
	if x := -14; x > 0 {
		resultMessage = "More Than zero"
	} else {
		resultMessage = "Less Than zero"
	}

	fmt.Println(resultMessage)

}
