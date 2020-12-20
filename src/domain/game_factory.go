package domain

type MinesweeperGameFactory interface {
	CreateGame(rowCount, colCount int, bombCount int) (*Game, error)
}

type minesweeperGameFactory struct {
	bombLocator BombLocator
}

func NewMinesweeperGameFactory(locator BombLocator) MinesweeperGameFactory {
	return &minesweeperGameFactory{bombLocator: locator}
}

func (this *minesweeperGameFactory) CreateGame(rowCount, colCount int, bombCount int) (*Game, error){


	return nil, nil
}


