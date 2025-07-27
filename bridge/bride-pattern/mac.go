package bridepattern

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Print("printing the file through mac")
	m.printer.PrintFile()
}

func (m *Mac)SetPrinter(printer Printer){
	fmt.Println("setting the printer")
	m.printer = printer
}