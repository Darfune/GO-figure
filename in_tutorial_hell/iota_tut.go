package main

import (
	"fmt"
)



func main(){
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1,c2,c3)

	const (
		c4 = iota
		c5 
		c6 
	)
	fmt.Println(c4,c5,c6)


	const(
		a = (iota * 2) + 1
		b
		c
	)
	fmt.Println(a,b,c)

	// x = -2, y = -4, z = -5
	const(
		x = -(iota + 2)
		_
		y
		z
	)
	fmt.Println(x,y,z)
		
}
