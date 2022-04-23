package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Stdin = standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("You entered:", input)

}
