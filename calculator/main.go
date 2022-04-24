package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// circleRadius := 15.5
	// circumference := circleRadius * 2 * math.Pi
	// fmt.Println("Circumference is : ", circumference)

	// fmt.Printf("Formated Circumference is : %.2f\n", circumference)

	//Value 1
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	value1str, err1 := reader.ReadString('\n')

	if err1 != nil {
		panic("Value 1 should not be empty!")
	}

	value1, err1 := strconv.ParseFloat(strings.TrimSpace(value1str), 64)

	if err1 != nil {
		panic("Value 1 must be a Number!")
	}

	fmt.Print("Value 2: ")
	value2str, err1 := reader.ReadString('\n')

	if err1 != nil {
		panic("Value 2 should not be empty!")
	}

	value2, err1 := strconv.ParseFloat(strings.TrimSpace(value2str), 64)

	if err1 != nil {
		panic("Value 2 must be a Number!")
	}

	sumCalc := value1 + value2

	//Missed
	sumCalc = math.Round(sumCalc*100) / 100

	fmt.Printf("The sum of %v and %v is %v", value1, value2, sumCalc)
}
