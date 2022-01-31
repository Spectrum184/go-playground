package adapter

import "fmt"

type Window struct{

}

func (w *Window) insertSquarePort()  {
	fmt.Println("Insert square port into window!")
}