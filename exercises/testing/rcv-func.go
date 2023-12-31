//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health, maxHealth uint
	energy, maxEnergy uint
	name              string
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Healed", amount, "health ->", player.health)
}

func (player *Player) applyDamage(amount uint) {
	if player.health < amount {
		player.health = 0
		fmt.Println(player.name, "died")
	} else {
		player.health -= amount
		fmt.Println(player.name, "Took damage", amount, "health ->", player.health)
	}
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println(player.name, "Gained energy", amount, "energy ->", player.energy)
}

func (player *Player) reduceEnergy(amount uint) {
	if player.energy < amount {
		player.energy = 0
		fmt.Println(player.name, "is exhausted")
	} else {
		player.energy -= amount
		fmt.Println(player.name, "Lost energy", amount, "energy ->", player.energy)
	}
}

func main() {
	player := Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.applyDamage(99)
	player.addHealth(10)
	player.reduceEnergy(450)
	player.addEnergy(10)

	player.reduceEnergy(450)
	player.applyDamage(30)
}
