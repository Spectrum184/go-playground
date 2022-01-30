package state

import "fmt"

type Machine struct{
	current state
}

func NewMachine() *Machine  {
	fmt.Printf("Machine is ready \n")
	return &Machine{NewOff()}
}

func (m *Machine) setCurrentSate(s state){
	m.current = s
}

func (m *Machine) On(){
	m.current.on(m)
}

func (m *Machine) Off(){
	m.current.off(m)
}