package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello World")
	
	unused := 100
	_ = unused

	// comment

	/*
	multiple line comment
	*/

	// var mv int //max value
	// var _ = mv
	// var packetsReceived int // Too Long name
	// _ = packetsReceived
	//
	// const MAX_VALUE = 100 // not ideal
	// const N = 100 //ideal

	// maxValue := 1000 // recommended
	// max_value := 1000 // not recommended

	name := "John Doe"
	fmt.Println("Hello, my name is ", name)

	a, b := 6, 2
	a, b = b, a
	a, c := 4, 6 // can't use := if there is not a new variable
	_ = c
	fmt.Println("Sum of a and b:", a+b, "| Mean Value of a and b:", (a+b)/2)

	fmt.Println("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f\n", 5, 6.8)
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi
	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The diameter of a %s with a Radius of %d si %f\n", figure, radius, float64(radius)*2*pi)

	//%q for qouted string
	fmt.Printf("This is a %q\n", figure)

	//%v can be used for any type of variable
	fmt.Printf("The diameter of a %v with a Radius of %v si %v\n", figure, radius, float64(radius)*2*pi)

	//%T to output the type of the argument
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	//%t for boolean
	close := true
	fmt.Printf("File closed: %t\n", close)

	//%b turns argument to base2
	fmt.Printf("%b\n", radius)
	fmt.Printf("%b\n", 55)
	fmt.Printf("%b\n", 100)
	fmt.Printf("%b\n", "Hello world")
	fmt.Printf("%b\n", figure)
	fmt.Printf("%b\n", pi)

	fmt.Printf("%08b\n", 45) // print number in 8 bit (usefull for ipv4 andresses)

	// turn decimal to hex decima number
	fmt.Printf("%x\n", 100)

	x := 3.5
	y := 6.9
	fmt.Printf("x * y = %.3f\n", x*y)
}
