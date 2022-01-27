package factory

type ISock interface{
	setLogo(logo string)
	setSize(size int)

	GetLogo() string
	GetSize() int
}

type sock struct{
	logo string
	size int
}

func (s *sock) setLogo(logo string)  {
	s.logo = logo
}

func (s *sock) setSize(size int){
	s.size = size
}

func (s *sock) GetLogo() string{
	return s.logo
}

func (s *sock) GetSize() int {
	return s.size
}