package server

import (
	"github.com/minesweeper/src/common/factory"
	"github.com/minesweeper/src/controllers"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/services"
)


func RegisterFactories( ) {
	//service
	factory.FactoryRegistrySingleton().RegisterFactory("default_MinesweeperService", services.CreateMinesweeperService)

	//mines locators
	factory.FactoryRegistrySingleton().RegisterFactory("random_MineLocator", domain.CreateRandomMinesLocator)
	factory.FactoryRegistrySingleton().RegisterFactory("fixed_MineLocator", domain.CreateFixedMinesLocator)

	//dals
	factory.FactoryRegistrySingleton().RegisterFactory("inmemory_GameDal", gamedal.CreateInMemoryGameDal)
	factory.FactoryRegistrySingleton().RegisterFactory("db_GameDal", gamedal.CreateDbGameDal)
}

func CreateController() *controllers.MinesweeperController {
	return &controllers.MinesweeperController{}
}