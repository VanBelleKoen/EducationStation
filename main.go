package main

import (
	TafelsTot10 "educationStation/multiplication"
	"fmt"
)

func main() {
	var i int
	fmt.Println("Wat zou je graag willen leren ?")
	fmt.Println("1. De tafels")
	fmt.Scan(&i)
	if i == 1 {
		TafelsTot10.Tafels()
	}
}
