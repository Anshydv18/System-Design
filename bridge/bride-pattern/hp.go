package bridepattern

import "fmt"

type HP struct {
}

func (pr *HP) PrintFile() {
	fmt.Println("printing through hp printer")
}