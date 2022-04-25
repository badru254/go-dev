package main

import (
	"fmt"
)

func main() {

	colors := []string{"Red", "Green", "Blue"}

	fmt.Println(colors)

	//Normal for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	//Shortened For loop
	for i := range colors {
		fmt.Println(colors[i])
	}

	//For Each loop
	for i, v := range colors {
		fmt.Printf("%v = %v\n", i, v)
	}

	//For Each loop with ignored index
	for _, v := range colors {
		fmt.Printf("The color is %v\n", v)
	}

	//While loop using For
	value := 1

	for value < 10 {
		fmt.Println(value)
		value++
	}

	//Break : jump to the end of the code block ---- works with (for & switch)
	//Continue : start back at the begining of code block ---- works with (for & switch)
	// Goto : jump to the end of the code block ---- works with (for & switch)
	sum := 1

	for sum < 1000 {
		sum += sum

		fmt.Println(sum)

		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of Program")

}
