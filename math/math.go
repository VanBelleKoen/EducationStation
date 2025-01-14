package math

import (
	Multiplication "educationStation/math/tables"
	fundamentals "educationStation/support"
	"fmt"
)

func Picker() {
	fmt.Println("Wat zou je graag willen leren ? \n 1. De tafels")
	result, _ := fundamentals.FetchInput()
	switch result.IntValue {
	case 1:
		Multiplication.Exercises()
	}
}