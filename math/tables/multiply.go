package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
)

func Multiplication() bool {
	left := (rand.Intn(max-min+1) + min)
	right := (rand.Intn(max-min+1) + min)

	fmt.Printf("\n %d x %d = ", left, right)
	result, _ := fundamentals.FetchInput()

	correct := left * right 
	if result.IntValue == correct {
		fmt.Println("Goed gedaan!")
		return true
	} else {
		fmt.Printf("Fout! Het juiste antwoord is %d\n", correct)
	}
	return false
}
