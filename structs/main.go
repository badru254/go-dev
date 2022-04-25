package main

import (
	"fmt"
)

func main() {
	mazda := Car{"Mazda", "Axela", "Automatic", 5}

	fmt.Println(mazda)
	fmt.Printf("%+v\n", mazda)
	fmt.Printf("Make :%v\nModel :%v\nTransmission :%v\nDoors :%v", mazda.Make, mazda.Model, mazda.Transmission, mazda.Doors)

	mazda.Model = "CX-5"
	fmt.Printf("Make :%v\nModel :%v\nTransmission :%v\nDoors :%v", mazda.Make, mazda.Model, mazda.Transmission, mazda.Doors)

}

//Car is a struct
type Car struct {
	Make         string
	Model        string
	Transmission string
	Doors        int
}
