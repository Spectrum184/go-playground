package bridge

type Computer interface{
	print()
	setPrinter(Printer)
}