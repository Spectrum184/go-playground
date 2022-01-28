package prototype

import "fmt"

type Folder struct{
	Children []INode
	Name string
}

func (f *Folder) Print(s string){
	fmt.Println(s + f.Name)

	for _, i := range f.Children{
		i.Print(s + s)
	}
}

func (f *Folder) Clone() INode{
	cloneFolder := &Folder{Name: f.Name + "_Clone"}

	var tempChildren []INode

	for _, i := range f.Children{
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.Children = tempChildren

	return cloneFolder
}


