package services

import (
	"github.com/minesweeper/src/common/apierrors"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/shared"
)

type MinesweeperService interface {
	GetGame(userId string, gameId string) (*shared.GameData, error)
	GetGameListByUserId(userId string) ([]*shared.GameData, error)

	NewGame(userId string, row, col int, minesCount int) (*shared.GameData, error)
	RevealCell(userId string, gameId string, row, col int) (*shared.GameData, error)
	MarkCell(userId string, gameId string, row, col int, mark shared.CellMarkType) (*shared.GameData, error)

	DeleteGame(userId string, gameId string) (error)
	DeleteAllGames(userId string) error
}


type minesweeperService struct {
	gameDal     gamedal.GameDal
	gameFactory domain.MinesweeperGameFactory
}


func (this *minesweeperService) GetGame(userId string, gameId string) (*shared.GameData, error){
	if userId == "" {
		return nil, apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

	if gameId == "" {
		return nil, apierrors.NewBadRequest(nil, "The game id is mandatory")
	}

	//get game
	savedGame, err := this.gameDal.GetGameById(userId, gameId)
	if err != nil {
		return nil, err
	}

	data := savedGame.GetData()

	data.Board.HideNotRevealed()

	return &data, nil
}

func (this *minesweeperService) GetGameListByUserId(userId string) ([]*shared.GameData, error) {
	if userId == "" {
		return nil, apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

	savedGames, err := this.gameDal.GetGameListByUserId(userId)
	if err != nil {
		return nil, err
	}

	output := make ([]*shared.GameData, 0)

	for i:=0; i<len(savedGames); i++ {
		game := savedGames[i]
		data := game.GetData()
		data.Board.HideNotRevealed()

		output = append(output, &data)
	}

	return output, nil
}


func (this *minesweeperService) NewGame(userId string, row, col int, minesCount int) (*shared.GameData, error){
	if userId == "" {
		return nil, apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

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



func (this *minesweeperService) RevealCell(userId string, gameId string, row, col int) (*shared.GameData, error){
	if userId == "" {
		return nil, apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

	if gameId == "" {
		return nil, apierrors.NewBadRequest(nil, "The game id is mandatory")
	}

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
	if userId == "" {
		return nil, apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

	if gameId == "" {
		return nil, apierrors.NewBadRequest(nil, "The game id is mandatory")
	}

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

func (this *minesweeperService) DeleteGame(userId string, gameId string) error {
	if userId == "" {
		return apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

	if gameId == "" {
		return apierrors.NewBadRequest(nil, "The game id is mandatory")
	}

	err := this.gameDal.DeleteGame(userId, gameId)

	if err != nil {
		return err
	}

	return nil
}

func (this *minesweeperService) DeleteAllGames(userId string) error {
	if userId == "" {
		return apierrors.NewBadRequest(nil, "The user id is mandatory")
	}

	err := this.gameDal.DeleteAllGames(userId)

	if err != nil {
		return err
	}

	return nil
}