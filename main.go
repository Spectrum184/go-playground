package main

import b "playground/bridge"

func main(){
	hpPrinter := &b.HP{}
	epsonPrinter := &b.Epson{}

	macComputer := &b.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	windowComputer := &b.Window{}
	windowComputer.SetPrinter(hpPrinter)
	windowComputer.Print()

	windowComputer.SetPrinter(epsonPrinter)
	windowComputer.Print()
}

