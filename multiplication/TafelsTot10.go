package TafelsTot10

import (
	"fmt"
	"math/rand"
	"strconv"
)

const min = 1
const max = 10

var answer string

func Tafels() {
	var i int
	fmt.Println("We zullen de tafels tot 10 leren")
	fmt.Println("1. Wil je de tafels rustig leren?")
	fmt.Println("2. Wil je de snel tafels leren?")

	fmt.Scan(&i)

	fmt.Println("We zullen de tafels tot 10 leren")
	fmt.Println("Geef als antwoord 'q' om te stoppen")

	for answer != "q" {
		if i == 1 {
			TafelsKalm()
		} else if i == 2 {
			TafelsSnel()
		} else {
			fmt.Println("Ongeldige keuze")
		}
	}
}

func TafelsKalm() {
	fmt.Println("")
	var left int = (rand.Intn(max-min+1) + min)
	var right int = (rand.Intn(max-min+1) + min)

	fmt.Printf("%d x %d = ", left, right)

	fmt.Scan(&answer)

	var correct int = left * right
	fmt.Println("")
	if answer == "q" {
		fmt.Println("Tot ziens!")
	} else if ans, err := strconv.Atoi(answer); err == nil && ans == correct {
		fmt.Println("Goed gedaan!")
	} else {
		fmt.Printf("Fout! Het juiste antwoord is %d\n", correct)
	}
}

func TafelsSnel() {

}
