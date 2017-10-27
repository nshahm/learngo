package control

import (
	"fmt"
)

func FindDescriptionForOperation(operation string) {
	switch operation {
	case "+":
		fmt.Println("Sum operation")

	case "-":
		fmt.Println("Subtract operation")

	case "*":
		fmt.Println("Multiplication operation")

	case "/":
		fmt.Println("Division operation")
	default:
		fmt.Println("No operation found")
	}
}
