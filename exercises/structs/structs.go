//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinates struct {
	x, y int
}

type Rectangle struct {
	a, b Coordinates
}

func width(rect Rectangle) int {
	return rect.b.x -rect.a.x
}

func lenght(rect Rectangle) int {
	return rect.a.y -rect.b.y
}

func area(rect Rectangle) {
	fmt.Println("The area is", lenght(rect) * width(rect))
}

func perimeter(rect Rectangle) {
	fmt.Println("The perimeter is", (width(rect) * 2) + (lenght(rect) * 2))
}

func main() {
	rect := Rectangle{a: Coordinates{0,7}, b: Coordinates{10,0}}
	area(rect)
	perimeter(rect)
	rect.a.y *= 2
	rect.b.x *= 2
area(rect)
	perimeter(rect)
}
