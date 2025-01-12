package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
	"strconv"
)

func Multiplication() {
	left := (rand.Intn(max-min+1) + min)
	right := (rand.Intn(max-min+1) + min)

	fmt.Printf("\n %d x %d = ", left, right)

	fmt.Scan(&answer)

	correct := left * right
	if answer == "q" {
		fundamentals.Quit()
	} else if ans, err := strconv.Atoi(answer); err == nil && ans == correct {
		fmt.Println("Goed gedaan!")
	} else {
		fmt.Printf("Fout! Het juiste antwoord is %d\n", correct)
	}
}
