package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I have done this tutorial at: ", n)

	t := time.Date(2021, time.August, 26, 13, 0, 0, 0, time.UTC)
	fmt.Println("I at lunch at: ", t)

	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Thu Aug 26 13:00:00 2021")

	//if err == nil {
	fmt.Printf("The type if parseTime is %T\n", parsedTime)
	//}

}
