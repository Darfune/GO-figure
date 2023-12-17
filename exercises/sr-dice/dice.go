//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	var diceNum = 2

	for i := 1; i <= 10; i++ {
		var total = 0
		for dice := 1; dice <= diceNum; dice++ {
			total += rng.Intn(7)
		}
		switch {
		case total == 2 && diceNum == 2:
			fmt.Print("Snake eyes")
		case total == 7:
			fmt.Print("Lucky 7")

		case total%2 == 0:
			fmt.Print("Even")
		case total%2 != 0:
			fmt.Print("Odd")
		}
		fmt.Println("->", total)
	}
	

}
