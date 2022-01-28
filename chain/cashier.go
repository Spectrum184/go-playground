package chain

import "fmt"

type Cashier struct{
	next Department
}

func (c *Cashier) Execute(p *Patient){
	if p.isPaid{
		fmt.Println("Patient is already paid their bill")

		return
	}

	fmt.Println("Patient is paiding")

	p.isPaid = true
}

func (c *Cashier) SetNext(next Department)  {
	c.next = next
}