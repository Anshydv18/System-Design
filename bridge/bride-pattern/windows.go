package bridepattern

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("printing through windows")
	w.printer.PrintFile()
}

func (m *Windows)SetPrinter(printer Printer){
	fmt.Println("setting the printer")
	m.printer = printer
}