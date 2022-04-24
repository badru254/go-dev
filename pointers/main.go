package main

import "fmt"

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p: ", *p)

	value1 := 42.34
	pointer1 := &value1

	fmt.Println("Value of pointer1: ", *pointer1)

	*pointer1 = 23.5 * *pointer1

	fmt.Println("Value of pointer1 multiplied by 23.5: ", *pointer1)

	fmt.Println("Value of value1 when pointer1 is multiplied by 23.5: ", value1)

}
