package support

import (
	"fmt"
)

func Quit() {
	fmt.Println("Tot ziens!")
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

