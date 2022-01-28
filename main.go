package main

import c "playground/chain"

func main(){
	cashier := &c.Cashier{}

	medical := &c.Medical{}
	medical.SetNext(cashier)

	doctor := &c.Doctor{}
	doctor.SetNext(medical)

	reception := &c.Reception{}
	reception.SetNext(doctor)

	patient := &c.Patient{Name: "Thanh"}

	reception.Execute(patient)

}

