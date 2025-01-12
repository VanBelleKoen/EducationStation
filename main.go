package main

import (
	math "educationStation/math"
	"fmt"
)

func main() {
	var i int
	fmt.Println("Wat zou je graag willen leren ?")
	fmt.Println("1. Wiskunde")
	fmt.Scan(&i)
	if i == 1 {
		math.Picker()
	}
}
