package composite

import "fmt"

type File struct{
	Name string
}

func (f *File) Search(keyword string)  {
	fmt.Printf("Searching for keyword %s \n" , keyword)
}