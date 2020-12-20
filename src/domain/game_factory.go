package domain

import (
	"errors"
)

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

	if rowCount == 0 || colCount == 0 {
		return nil, errors.New("Row and Column size can not be 0 or less")
	}

	if bombCount == 0 {
		return nil, errors.New("The count of bombs must be greater than 0")
	}

	if rowCount * bombCount <= bombCount {
		return nil, errors.New("The count of bombs can not be greater than the count of cells")
	}

	//create new game
	game, err := NewGame(rowCount, colCount, bombCount, this.bombLocator)

	return &game, err
}


