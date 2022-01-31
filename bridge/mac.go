package bridge

import "fmt"

type Mac struct{
	printer Printer
}

func (w *Mac) Print()  {
	fmt.Println("Print request for mac")
	w.printer.printFile()
}

func (w *Mac) SetPrinter(printer Printer)  {
	w.printer = printer
}