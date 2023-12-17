//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func twoTwos() (int, int) {
	return 2, 2
}

func five() int {
	return 5
}

func addThree(a, b, c int) int {
	return a + b + c
}

func greetSomeone(name string) {
	fmt.Println("Hello", name)
}

func hiThere() string {
	return "hi"
}

func main() {
	greetSomeone("Alice")
	fmt.Println(hiThere())
	a, b :=  twoTwos()
	answer := addThree(five(), a, b)
	fmt.Println(answer)
}
