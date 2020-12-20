package services

import (
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
)

type MinesweeperService interface {

}


type minesweeperService struct {
	gameDal     gamedal.GameDal
	gameFactory domain.MinesweeperGameFactory
}



