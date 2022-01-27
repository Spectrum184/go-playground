package factory

type Nike struct{}

func (n *Nike) MakeShoe() IShoe  {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *Nike) MakeSock() ISock{
	return &nikeSock{
		sock: sock{
			logo: "nike",
			size: 20,
		},
	}
}