package main

import (
	"fmt"
)

//In go, a method is a member of a type
func main() {
	mazda := Car{"Mazda", "Axela", "Automatic", 5}

	fmt.Println(mazda)
	fmt.Printf("%+v\n", mazda)
	fmt.Printf("Make :%v\nModel :%v\nTransmission :%v\nDoors :%v\n", mazda.Make, mazda.Model, mazda.Transmission, mazda.Doors)
	fmt.Println("---Body Type----")
	mazda.Bodytype()
	mazda.Model = "CX-5"
	fmt.Printf("Make :%v\nModel :%v\nTransmission :%v\nDoors :%v\n", mazda.Make, mazda.Model, mazda.Transmission, mazda.Doors)
	fmt.Println("---Body Type----")
	mazda.Bodytype()

}

//Car is a struct
type Car struct {
	Make         string
	Model        string
	Transmission string
	Doors        int
}

//Bodytype is the body type of the car
func (c Car) Bodytype() {
	switch c.Model {
	case "CX-5":
		fmt.Println("Cross over SUV")
	case "CX-3":
		fmt.Println("Cross over SUV")
	case "CX-30":
		fmt.Println("Cross over SUV")
	case "Axela":
		fmt.Println("Estate")
	case "Atenza":
		fmt.Println("Estate")
	case "CX-7":
		fmt.Println("SUV")
	case "CX-9":
		fmt.Println("SUV")
	case "CX-90":
		fmt.Println("SUV")
	default:
		fmt.Println("Body Type Unknown")
	}
}
