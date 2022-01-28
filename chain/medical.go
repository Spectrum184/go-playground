package chain

import "fmt"

type Medical struct{
	next Department
}

func (m *Medical) Execute(p *Patient){
	if p.isMedicineProvied{
		fmt.Println("Patient is already provied medicine")

		m.next.Execute(p)

		return
	}

	fmt.Println("We are providing to patient")

	p.isMedicineProvied = true

	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department)  {
	m.next = next
}