package main

import (
	"fmt"
	"math"
)

func main() {

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Println("Circumference is : ", circumference)

	fmt.Printf("Formated Circumference is : %.2f\n", circumference)

}
