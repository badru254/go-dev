package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = []string{"Red", "Blue", "Green"}

	fmt.Println(colors)

	//adding
	colors = append(colors, "Maroon")
	fmt.Println(colors)

	//deleting
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)

	numbers[0] = 132
	numbers[1] = 122
	numbers[2] = 344
	numbers[3] = 3
	numbers[4] = 23
	fmt.Println(numbers)
	numbers = append(numbers, 34)

	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
