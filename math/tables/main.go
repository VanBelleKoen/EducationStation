package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
)

const (
	min = 1
	max = 10
)

var answer string

func Multiply() {
	var choice int
	fmt.Println("We zullen de tafels tot 10 leren. \n 1. Wil je de tafels rustig leren? \n 2. Wil je de tafels op snelheid leren?")

	fmt.Scan(&choice)

	fmt.Println("We zullen de tafels tot 10 leren \n Geef als antwoord 'q' om te stoppen")

	for answer != "q" {
		switch choice {
		case 1:
			WithoutTimeLimit()
		case 2:
			fundamentals.SetTimeLimit()
			Multiplication()
		default:
			fmt.Println("Ongeldige keuze")
		}
	}
}
	
	func WithoutTimeLimit() {
		for answer != "q" {
			switch rand.Intn(2) {
			case 0:
				Multiplication()
			case 1:
				Division()
			}
		}
	}