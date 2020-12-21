package domain

import (
	"github.com/minesweeper/src/common/apierrors"
)

type MinesweeperGameFactory interface {
	CreateGame(rowCount, colCount int, minesCount int) (*Game, error)
}

type minesweeperGameFactory struct {
	minesLocator MinesLocator
}

func NewMinesweeperGameFactory(locator MinesLocator) MinesweeperGameFactory {
	return &minesweeperGameFactory{minesLocator: locator}
}

func (this *minesweeperGameFactory) CreateGame(rowCount, colCount int, minesCount int) (*Game, error){

	if rowCount == 0 || colCount == 0 {
		return nil, apierrors.NewBadRequest(nil, "Row and Column size can not be 0 or less")
	}

	if minesCount == 0 {
		return nil, apierrors.NewBadRequest(nil, "The count of mines must be greater than 0")
	}

	if rowCount *minesCount <= minesCount {
		return nil, apierrors.NewBadRequest(nil, "The count of mines can not be greater than the count of cells")
	}

	//create new game
	game, err := NewGame(rowCount, colCount, minesCount, this.minesLocator)

	return &game, err
}


