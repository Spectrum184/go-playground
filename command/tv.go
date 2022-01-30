package command

import "fmt"

type TV struct{
	isRunning bool
}

func (tv *TV) on(){
	tv.isRunning = true
	fmt.Println("TV is running")
}

func (tv *TV) off(){
	tv.isRunning = false

	fmt.Println("TV off")
}