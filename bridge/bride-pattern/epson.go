package bridepattern

import "fmt"

type Epson struct {
}

func (ep *Epson) PrintFile() {
	fmt.Println("printing through epson printer")
}
