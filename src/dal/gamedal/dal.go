package gamedal

import "github.com/minesweeper/src/domain"

type GameDal interface {
	//Both key to validate that de game belongs to the user
	GetGameById(userId string, gameId string) (*domain.Game, error)
	GetGameListByUserId(userId string) ([]*domain.Game, error)
	InsertGame(userId string, game *domain.Game) (*domain.Game, error)
	UpdateGame(game *domain.Game) (*domain.Game, error)

	DeleteGame(userId string, gameId string) error
	DeleteAllGames(userId string) error
}