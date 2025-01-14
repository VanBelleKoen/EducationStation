package support

import (
	"fmt"
	"os"
	"strconv"
)

func Quit() {
	fmt.Println("Tot ziens!")
	os.Exit(0)
}

func SetTimeLimit() {
	var timeLimit int
	fmt.Println("Hoeveel minuten wil je voor deze oefeningen?")
	fmt.Scan(&timeLimit)

	fmt.Println("Je hebt", timeLimit, "minuten om deze oefeningen te maken.")
}

func TimeIsUp() {
	fmt.Println("De tijd is op!")
}

func Correct() {
	fmt.Println("Goed gedaan!")
}

type InputResult struct {
	IntValue   int
	StringValue string
	IsInt      bool
}

func FetchInput() (InputResult, error) {
	var input string
	fmt.Scan(&input)
	if input == "q" {
			Quit()
			return InputResult{}, fmt.Errorf("quit signal received")
	}

	if value, err := strconv.Atoi(input); err == nil {
			return InputResult{IntValue: value, IsInt: true}, nil
	} else {
			return InputResult{StringValue: input, IsInt: false}, nil
	}
}

