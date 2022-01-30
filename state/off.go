package state

import "fmt"

type off struct{}


func NewOff() state{
	return &off{}
}

func (o *off) on(m *Machine){
	fmt.Println("Machine already ON")
}

func (o *off) off(m *Machine){
	fmt.Println("Machine is going OFF")
	m.setCurrentSate(NewOn())
}