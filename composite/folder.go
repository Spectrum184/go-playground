package composite

import "fmt"

type Folder struct{
	components []Component
	Name string
}

func (f *Folder) Search(keyword string){
	fmt.Printf("Searching recursively for key %s \n", keyword)

	for _, composite := range f.components{
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component){
	f.components = append(f.components, c)
}