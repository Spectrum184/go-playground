package factory

type Adidas struct{

}

func (a *Adidas) MakeShoe() IShoe{
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeSock() ISock{
	return &adidasSock{
		sock: sock{
			logo: "adidas",
			size: 15,
		},
	}
}
