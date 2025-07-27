package main

import (
	factorydesign "factory/factory_design"
	"fmt"
)

func main() {
	gun1 := factorydesign.GetGun("AK47")

	fmt.Println("Gun Name:", gun1.GetName())
	fmt.Println("Gun Power:", gun1.GetPower())
}
