package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type DiceRoll int
type PartialRoll int

var (
	randomRollGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func CalculateRoll(a, b PartialRoll) DiceRoll{
	return DiceRoll(a ^ b)
}

func RollPartially() PartialRoll {
	
	roll := randomRollGenerator.Int31n(6)
	for roll == 0 {
		roll = randomRollGenerator.Int31n(6)
	}

	return PartialRoll(roll)
}

func PrintDiceRollWinner(client, server DiceRoll) {

	if client > server {
		
		fmt.Println("Alice Won!")
	} else if client < server {
		
		fmt.Println("Bob Won!")
	} else {

		fmt.Println("Draw!")
	}
}