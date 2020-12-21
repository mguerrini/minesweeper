package gamedal

import (
	"github.com/minesweeper/src/common/configs"
	"github.com/minesweeper/src/domain"
)

type GameDbDal struct {
	connectionString string
}


func CreateDbGameDal(configurationName string) (interface{}, error){
	cnn, err := configs.Singleton().GetString(configurationName + ".connectionstring")

	if err != nil {
		return nil, err
	}
	return &GameDbDal{connectionString: cnn}, nil
}


func (g GameDbDal) GetGameById(userId, gameId string) (*domain.Game, error) {
	panic("implement me")
}

func (g GameDbDal) GetGameListByUserId(userId string) ([]domain.Game, error) {
	panic("implement me")
}

func (g GameDbDal) InsertGame(userId string, game *domain.Game) (domain.Game, error) {
	panic("implement me")
}

func (g GameDbDal) UpdateGame(game *domain.Game) (domain.Game, error) {
	panic("implement me")
}

func NewDbGameDal(factoryConfigurationName string) (GameDal, error) {
	return &GameDbDal{}, nil
}



