package adapter

import "fmt"

type Mac struct{}

func (m *Mac) insertCirclePort()  {
	fmt.Println("Insert circle port into Mac!")
}