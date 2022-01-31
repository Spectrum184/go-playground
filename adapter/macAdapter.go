package adapter

type MacAdapter struct{
	M Mac
}

func (macAdapter *MacAdapter) insertSquarePort()  {
	macAdapter.M.insertCirclePort()
}