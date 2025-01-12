package math

import (
	Multiplication "educationStation/math/tables"
	"fmt"
)

func Picker() {
	var choice int
	fmt.Println("Wat zou je graag willen leren ? \n 1. De tafels")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		Multiplication.Exercises()
	}
}