package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
	"strconv"
)

func Multiplication() bool {
	left := (rand.Intn(max-min+1) + min)
	right := (rand.Intn(max-min+1) + min)

	fmt.Printf("\n %d x %d = ", left, right)
	fmt.Scan(&Answer)

	correct := left * right
	if Answer == "q" {
		fundamentals.Quit()
	} else if ans, err := strconv.Atoi(Answer); err == nil && ans == correct {
		fmt.Println("Goed gedaan!")
		return true
	} else {
		fmt.Printf("Fout! Het juiste antwoord is %d\n", correct)
	}
	return false
}
