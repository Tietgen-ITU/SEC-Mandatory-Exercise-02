package utils

import "fmt"

type DiceRoll int
type PartialRoll int

func CalculateRoll(a, b PartialRoll) DiceRoll{
	return DiceRoll(a ^ b)
}

func PrintDiceRollWinner(client, server DiceRoll) {

	if client > server {
		
		fmt.Println("Client Won!")
	} else if client < server {
		
		fmt.Println("Server Won!")
	} else {

		fmt.Println("Draw!")
	}
}