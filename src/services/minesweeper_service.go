package services

import (
	"errors"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/shared"
)

type MinesweeperService interface {

}


type minesweeperService struct {
	gameDal     gamedal.GameDal
	gameFactory domain.MinesweeperGameFactory
}



func (this *minesweeperService) NewGame(userId string, row, col int, bombsCount int) (*shared.GameData, error){
	game, err := this.gameFactory.CreateGame(row, col, bombsCount)

	if err != nil {
		return nil, err
	}

	//save it
	savedGame := this.gameDal.InsertGame(userId, game)

	//get data of the game
	data := savedGame.GetData()

	data.Board.HideNotRevealed()

	return &data, nil
}


func (this *minesweeperService) RevealCell(userId string, gameId string, row, col int) (*shared.GameData, error){

	//save it
	savedGame := this.gameDal.GetGameById(userId, gameId)

	if savedGame == nil {
		return nil, errors.New("The gameid is invalid")
	}

	err := savedGame.RevealCell(row, col)

	if err != nil {
		return nil, err
	}

	//update the game
	this.gameDal.UpdateGame(savedGame)

	//get data of the game
	data := savedGame.GetData()

	if savedGame.IsFinished() {
		data.Board.RevealAll()
	} else {
		data.Board.HideNotRevealed()
	}

	return &data, nil
}


func (this *minesweeperService) MarkCell(userId string, gameId string, row, col int, mark shared.CellMarkType) (*shared.GameData, error){

	//save it
	savedGame := this.gameDal.GetGameById(userId, gameId)

	if savedGame == nil {
		return nil, errors.New("The gameid is invalid")
	}

	err := savedGame.MarkCell(row, col, mark)

	if err != nil {
		return nil, err
	}

	//update the game
	this.gameDal.UpdateGame(savedGame)

	//get data of the game
	data := savedGame.GetData()

	data.Board.HideNotRevealed()

	return &data, nil
}

