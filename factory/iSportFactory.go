package factory

type ISportFactory interface{
	MakeShoe() IShoe
	MakeSock() ISock
}

func GetSportFactory(brand string) ISportFactory  {
	switch brand {
	case "adidas":
		 return &Adidas{}

	case "nike":
		return &Nike{}
		}

		return nil

}