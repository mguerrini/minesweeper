package gamedal

import "github.com/minesweeper/src/domain"

type GameDal interface {
	//Both key to validate that de game belongs to the user
	GetGameById(userId string, gameId string) *domain.Game
	GetGameListByUserId(userId string) []domain.Game
	InsertGame(userId string, game *domain.Game) domain.Game
	UpdateGame(game *domain.Game) domain.Game
}