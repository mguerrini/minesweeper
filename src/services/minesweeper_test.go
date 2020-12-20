package services

import (
	"github.com/minesweeper/src/common/configs"
	"github.com/minesweeper/src/common/factory"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"testing"
)

func servicesTestSetup() {
	//service
	factory.FactoryRegistrySingleton().RegisterFactory("default_MinesweeperService", CreateMinesweeperService)

	//bomb locators
	factory.FactoryRegistrySingleton().RegisterFactory("random_BombLocator", domain.CreateRandomBombLocator)
	factory.FactoryRegistrySingleton().RegisterFactory("fixed_BombLocator", domain.CreateFixedBombLocator)

	//dals
	factory.FactoryRegistrySingleton().RegisterFactory("inmemory_GameDal", gamedal.CreateInMemoryGameDal)
	factory.FactoryRegistrySingleton().RegisterFactory("db_GameDal", gamedal.CreateDbGameDal)
}

func Test_CreateMinesweeperSingleton_FromConfiguration1(t *testing.T) {
	servicesTestSetup()

	service, err := NewMinesweeperService("root.services.minesweeper.default")

	if err != nil {
		t.Error("Error creating Minesweeper Service: ", err)
	}

	validateServiceWithDbDal(t, service)
}

func Test_CreateMinesweeperSingleton_FromConfiguration2(t *testing.T) {
	servicesTestSetup()

	service, err := NewMinesweeperService("root.services.minesweeper.local")

	if err != nil {
		t.Error("Error creating Minesweeper Service: ", err)
	}
	validateServiceWithInMemoryDal(t, service)
}

//Test creating passing empty configuration name
func Test_CreateMinesweeperSingleton_Default1(t *testing.T) {
	servicesTestSetup()

	//try to build service with configuration name
	service1, err1 := NewMinesweeperService("")
	if err1 != nil {
		t.Error("Error creating Minesweeper Service: ", err1)
	}
	validateServiceWithDbDal(t, service1)
}

//Test creating with not existent configuration
func Test_CreateMinesweeperSingleton_Default2(t *testing.T) {
	servicesTestSetup()

	//clean the configuration
	configs.Singleton().Clean()

	//try to build service with configuration name
	service1, err1 := NewMinesweeperService("root.services.minesweeper.default")
	if err1 != nil {
		t.Error("Error creating Minesweeper Service: ", err1)
	}
	validateServiceWithDbDal(t, service1)
}



func validateServiceWithDbDal(t *testing.T, service MinesweeperService) {

	if service == nil {
		t.Error("Error creating Minesweeper Service: Service nil")
	}

	//validate that is created with the correct components
	realType, ok := service.(*minesweeperService)

	if !ok {
		t.Error("The Minesweeper Service is not of type *minesweeperService")
	}

	_, ok = realType.gameDal.(*gamedal.GameDbDal)

	if !ok {
		t.Error("The GameDal of the Minesweeper Service is not of type *GameDbDal")
	}
}

func validateServiceWithInMemoryDal(t *testing.T, service MinesweeperService) {

	if service == nil {
		t.Error("Error creating Minesweeper Service: Service nil")
	}

	//validate that is created with the correct components
	realType, ok := service.(*minesweeperService)

	if !ok {
		t.Error("The Minesweeper Service is not of type *minesweeperService")
	}

	_, ok = realType.gameDal.(*gamedal.GameInMemoryDal)

	if !ok {
		t.Error("The GameDal of the Minesweeper Service is not of type *GameInMemoryDal")
	}
}
