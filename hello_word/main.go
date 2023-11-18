package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")
	
	// Print power of 2 in 11
	res := math.Pow(2,11);
	fmt.Println(res)

	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// Print 100 random numbers
	fmt.Println("")

	for i := 1; i <= 100; i++{
		rand.Seed(time.Now().UnixNano())
//		random := rand.New(rand.NewSourc))
		randomNumber := rand.Intn(100) + 1

		fmt.Printf("Random number: %d\n", randomNumber)
	}

	// Print current date and time
	currentTime := time.Now()
	fmt.Println("Current Date and Time:", currentTime)
}
