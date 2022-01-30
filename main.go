package main

import o "playground/observer"

func main(){
	shirtItem := o.NewItem("Nike shirt")
	observerFirst := &o.Customer{ID: "abc@gmail.com"}
	observerSecond := &o.Customer{ID: "xyz@hotmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.DeRegister(observerSecond)

	shirtItem.UpdateAvailability()

}

