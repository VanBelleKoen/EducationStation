package tables

import (
	fundamentals "educationStation/support"
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 10
)

var (
	Answer                 string
	AmountOfCorrectAnswers int
	AmountOfTotalAnswers   int
)

func Exercises() {
	fmt.Println("We zullen de tafels tot 10 leren. \n 1. Wil je de tafels rustig leren? \n 2. Wil je de tafels op snelheid leren?")

	result, _ := fundamentals.FetchInput()

	switch result.IntValue {
	case 1:
		for AmountOfTotalAnswers < 20 {
			Exercise()
		}
		fmt.Println("Je hebt", AmountOfCorrectAnswers, "van de", AmountOfTotalAnswers, "vragen goed beantwoord")
	case 2:
		fmt.Println("Hoeveel tijd wil je om de tafels te oefenen? \n 1. 30 seconden \n 2. 1 minuut \n 3. 2 minuten, \n 4. 5 minuten")
		result, _ := fundamentals.FetchInput()
		WithTimeLimit(result.IntValue)
	default:
		fmt.Println("Ongeldige keuze")
	}
}

func Exercise() {
	var correct bool
	switch rand.Intn(2) {
	case 0:
		correct = Multiplication()
	case 1:
		correct = Division()
	}
	if correct {
		AmountOfCorrectAnswers++
	}
	AmountOfTotalAnswers++
}

func WithTimeLimit(choiceOfTime int) {
	var StayFor int
	switch choiceOfTime {
	case 1:
		StayFor = 30
	case 2:
		StayFor = 60
	case 3:
		StayFor = 120
	case 4:
		StayFor = 300
	}

	timeout := time.After(time.Duration(StayFor) * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			fmt.Println("\nTime's up! \nJe hebt", AmountOfCorrectAnswers, "van de", AmountOfTotalAnswers, "vragen goed beantwoord")
			return
		case <-ticker.C:
			Exercise()
		}
	}
}
