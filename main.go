package main

import (
	math "educationStation/math"
	fundamentals "educationStation/support"
	"fmt"
)

func main() {
	fmt.Println("Type q in om te stoppen. \n Wat zou je graag willen leren ?")
	fmt.Println("1. Wiskunde")
	result, _ := fundamentals.FetchInput()
	if result.IntValue == 1 {
		math.Picker()
	}
}
