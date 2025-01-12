package main

import (
	math "educationStation/math"
	"fmt"
)

func main() {
	var choice int
	fmt.Println("Type q in om te stoppen. \n Wat zou je graag willen leren ?")
	fmt.Println("1. Wiskunde")
	fmt.Scan(&choice)
	if choice == 1 {
		math.Picker()
	}
}
