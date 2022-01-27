package main

import (
	"fmt"
	"playground/factory"
)

func main(){
	adidasFactory := factory.GetSportFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasSock := adidasFactory.MakeSock()

	printShoe(adidasShoe)
	printSock(adidasSock)

	nikeFactory := factory.GetSportFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeSock := nikeFactory.MakeSock()

	printShoe(nikeShoe)
	printSock(nikeSock)
}

func printShoe(s factory.IShoe)  {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printSock(s factory.ISock){
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}