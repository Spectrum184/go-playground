package bridge

import "fmt"

type Epson struct{}

func(e *Epson) printFile(){
	fmt.Println("Printing by epson printer")
}