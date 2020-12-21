package server

import (
	"github.com/minesweeper/src/common/factory"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/services"
)


func RegisterFactories( ) {
	//service
	factory.FactoryRegistrySingleton().RegisterFactory("default_MinesweeperService", services.CreateMinesweeperService)

	//mines locators
	factory.FactoryRegistrySingleton().RegisterFactory("random_MinesLocator", domain.CreateRandomMinesLocator)
	factory.FactoryRegistrySingleton().RegisterFactory("fixed_MinesLocator", domain.CreateFixedMinesLocator)

	//dals
	factory.FactoryRegistrySingleton().RegisterFactory("inmemory_GameDal", gamedal.CreateInMemoryGameDal)
	factory.FactoryRegistrySingleton().RegisterFactory("db_GameDal", gamedal.CreateDbGameDal)
}
