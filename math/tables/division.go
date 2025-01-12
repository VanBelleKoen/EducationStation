package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
	"strconv"
)

func Division() {
	left := (rand.Intn(max-min+1) + min)
	right := (rand.Intn(max-min+1) + min)

	bigNumber := left * right

	fmt.Printf("\n %d : %d = ", bigNumber, left)
	fmt.Scan(&answer)

	if answer == "q" {
		fundamentals.Quit()
	} else if ans, err := strconv.Atoi(answer); err == nil && ans == right {
		fmt.Println("Goed gedaan!")
	} else {
		fmt.Printf("Fout! Het juiste antwoord is %d\n", right)
	}
}
