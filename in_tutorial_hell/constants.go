package main

import (
	"fmt"
)

func main(){
	// constants don't have to be used unlike normal variables
	const days int = 7

	var i int
	fmt.Println(i)

	// const have to be initialized unlike normal varaible
	const pu float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hour
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	x, y := 5, 0
	_, _ = x, y

	// Run time error
	// fmt.Println(x/y)

	const a = 5
	const b = 0

	// compile time error
	// fmt.Println(a/b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)
	fmt.Println(min1, min2, min3)

	const (
		min4 = -500
		min5
		min6
	)
	fmt.Println(min4, min5, min6)
}
