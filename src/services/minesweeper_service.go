package services

import (
	"github.com/minesweeper/src/common/apierrors"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/shared"
)

type MinesweeperService interface {
	NewGame(userId string, row, col int, minesCount int) (*shared.GameData, error)
	RevealCell(userId string, gameId string, row, col int) (*shared.GameData, error)
	MarkCell(userId string, gameId string, row, col int, mark shared.CellMarkType) (*shared.GameData, error)
	GetGame(userId string, gameId string) (*shared.GameData, error)
}


type minesweeperService struct {
	gameDal     gamedal.GameDal
	gameFactory domain.MinesweeperGameFactory
}



func (this *minesweeperService) NewGame(userId string, row, col int, minesCount int) (*shared.GameData, error){
	game, err := this.gameFactory.CreateGame(row, col, minesCount)

	if err != nil {
		return nil, err
	}

	//save it
	savedGame, err := this.gameDal.InsertGame(userId, game)
	if err != nil {
		return nil, err
	}


	//get data of the game
	data := savedGame.GetData()

	data.Board.HideNotRevealed()

	return &data, nil
}

func (this *minesweeperService) GetGame(userId string, gameId string) (*shared.GameData, error){
	//get game
	savedGame, err := this.gameDal.GetGameById(userId, gameId)
	if err != nil {
		return nil, err
	}

	data := savedGame.GetData()

	data.Board.HideNotRevealed()

	return &data, nil
}


func (this *minesweeperService) RevealCell(userId string, gameId string, row, col int) (*shared.GameData, error){

	//get game
	savedGame, err := this.gameDal.GetGameById(userId, gameId)
	if err != nil {
		return nil, err
	}


	if savedGame == nil {
		return nil, apierrors.NewBadRequest(nil, "The gameid is invalid")
	}

	err = savedGame.RevealCell(row, col)

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
	savedGame, err := this.gameDal.GetGameById(userId, gameId)
	if err != nil {
		return nil, err
	}

	if savedGame == nil {
		return nil, apierrors.NewBadRequest(nil, "The gameid is invalid")
	}

	err = savedGame.MarkCell(row, col, mark)

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

