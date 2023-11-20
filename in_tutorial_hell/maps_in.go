package main

import (
	"log"
)

type Name struct{
	FirstName string
	LastName string
}

func main(){
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]Name)
	me := Name{
		FirstName: "Trevor",
		LastName: "Sawler",
	}

	myMap2["me"] = me

	log.Println(myMap2["me"].FirstName)
}
