package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}
/*
Private variables and funcs start with a lowercase letter,
while Public vars and funcs start with an Uppercase letter
*/
func main(){

	// var s2 = "six"
	// s := "eight"
	//
	// log.Println("s is", s)
	// log.Println("s2 is", s2)
	//
	// saySomething("xxx")

	user := User{
		FirstName: "Trevor",
		LastName: "Sawler",
	}

	log.Println(user.FirstName, user.LastName, "Birthday:", user.BirthDate)
}

func saySomething(s3 string) (string, string){
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
