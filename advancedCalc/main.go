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

	//Initialize Calculator with lables
	calculator := Calculator{Value1Label: "Value 1", Value2Label: "Value 2", Value3Label: "Select an operation (*,+,-,/)"}
	//Value 1
	_, calculator.Value1 = calculator.GetValue(1)
	//Value 2
	_, calculator.Value2 = calculator.GetValue(2)
	//Operation
	calculator.Value3, _ = calculator.GetValue(3)
	//Get Result
	calculator.Result = calculator.calculate()

	fmt.Println(calculator.Result)

}

type Calculator struct {
	Value1Label string
	Value2Label string
	Value3Label string
	Value1      float64
	Value2      float64
	Value3      string
	Result      string
}

//Validate get value from input
func (c Calculator) GetValue(valueIndex int) (string, float64) {
	res1 := ""
	var res2 float64 = 0.0

	label := c.Value3Label

	if valueIndex == 1 {
		label = c.Value1Label
	} else if valueIndex == 2 {
		label = c.Value2Label
	}

	reader := bufio.NewReader(os.Stdin)
	label = label + ": "
	fmt.Print(label)

	valueStr, _ := reader.ReadString('\n')

	if valueIndex == 3 {
		res1 = c.ValidateOperationInput(valueStr)
	} else {
		res2 = c.ValidateNumInput(valueStr)
	}

	return res1, res2
}

//Validate operation input
func (c Calculator) ValidateOperationInput(input string) string {
	if strings.TrimSpace(input) == "" {
		panic("Value not  be Empty!")
	}
	return strings.TrimSpace(input)
}

//Validate numerical input
func (c Calculator) ValidateNumInput(input string) float64 {
	value, err1 := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err1 != nil {
		panic("Value must be a Number!")
	}

	return value
}

//Do Calculations : Value1 operation Value2
func (c Calculator) calculate() string {
	result := 0.0
	operation := "multiplication"

	switch c.Value3 {
	case "+":
		result = c.Value1 + c.Value2
		operation = "sum"
	case "-":
		result = c.Value1 - c.Value2
		operation = "subtraction"
	case "/":
		result = c.Value1 / c.Value2
		operation = "division"
	default:
		result = c.Value1 * c.Value2
	}

	result = math.Round(result*100) / 100

	return fmt.Sprintf("The %v of %v and %v is %v", operation, c.Value1, c.Value2, result)

}
