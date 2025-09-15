package main

import (
	concreteobserver "observer/concrete_observer"
	concretesubjects "observer/concrete_subjects"
)

func main() {
	shirtItem := concretesubjects.NewItem("shirt")

	observerObj1 := &concreteobserver.Customer{Id: "ansh@gmail.com"}
	observerObj2 := &concreteobserver.Customer{Id: "ansh@kk.com"}

	shirtItem.Register(observerObj1)
	shirtItem.Register(observerObj2)

	shirtItem.NotifyObserver()
}
