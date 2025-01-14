package tables

import (
	"educationStation/support"
	"fmt"
	"math/rand"
)

func Division() bool {
    left := (rand.Intn(max-min+1) + min)
    right := (rand.Intn(max-min+1) + min)

    bigNumber := left * right

    fmt.Printf("\n %d : %d = ", bigNumber, left)
    result, err := support.FetchInput()
    if err != nil {
        fmt.Println(err)
        return false
    }

    if result.IsInt && result.IntValue == right {
        fmt.Println("Goed gedaan!")
        return true
    } else {
        fmt.Printf("Fout! Het juiste antwoord is %d\n", right)
        return false
    }
}