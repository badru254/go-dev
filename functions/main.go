package main

import (
	"fmt"
)

func main() {

	doSomething()

	sum := addValues(5, 8)

	fmt.Println("Sum: ", sum)

	sum2, valuesLen := addAlotOfValues(11, 22, 44, 55, 66, 77, 88, 99, 1010, 55)

	fmt.Printf("Sum of %v of Values: %v", valuesLen, sum2)
}

func doSomething() {
	fmt.Println("Doing Something!")
}

func addValues(value1, value2 int) int {

	return value1 + value2
}

func addAlotOfValues(nums ...int) (int, int) {

	total := 0

	for _, v := range nums {
		total += v
	}

	return total, len(nums)
}
