package main

import (
	bridepattern "bridge/bride-pattern"
	"fmt"
)

func main() {

	hpPrinter := &bridepattern.HP{}
	epsonPrinter := &bridepattern.Epson{}

	macComputer := &bridepattern.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &bridepattern.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
