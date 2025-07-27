package main

import (
	"cheese-pizza/classes"
	"fmt"
)

func main() {
	pizza := &classes.PizzaMania{}

	tommatoTopping := &classes.TommatoTopping{Pizza: pizza}
	tommatoprice := tommatoTopping.GetPrice()

	cheesetopped := &classes.CheeseToppingPizza{
		Pizza: tommatoTopping,
	}

	cheesedToppedPizaprice := cheesetopped.GetPrice()
	fmt.Println(tommatoprice)
	fmt.Println(cheesedToppedPizaprice)
}
