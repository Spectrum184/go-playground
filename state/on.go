package state

import "fmt"

type on struct{}

func NewOn() state{
	return &on{}
}

func (o *on) on(m *Machine){
	fmt.Println("Machine already ON")
}

func (o *on) off(m *Machine){
	fmt.Println("Machine is going OFF")
	m.setCurrentSate(NewOff())
}