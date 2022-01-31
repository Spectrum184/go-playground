package bridge

import "fmt"

type Window struct{
	printer Printer
}

func (w *Window) Print()  {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *Window) SetPrinter(printer Printer)  {
	w.printer = printer
}