package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
	"strconv"
)

func Division() bool {
	left := (rand.Intn(max-min+1) + min)
	right := (rand.Intn(max-min+1) + min)

	bigNumber := left * right

	fmt.Printf("\n %d : %d = ", bigNumber, left)
	fmt.Scan(&Answer)

	if Answer == "q" {
		fundamentals.Quit()
	} else if ans, err := strconv.Atoi(Answer); err == nil && ans == right {
		fmt.Println("Goed gedaan!")
		return true
	} else {
		fmt.Printf("Fout! Het juiste antwoord is %d\n", right)
	}
	return false
}
