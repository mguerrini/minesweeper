package services

import (
	"github.com/minesweeper/src/common/configs"
	"github.com/minesweeper/src/common/factory"
	"github.com/minesweeper/src/common/helpers"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/shared"
	"testing"
)

func servicesTestSetup() {
	configs.Initialize("testing/service_test_configuration.yml")

	//service
	factory.FactoryRegistrySingleton().RegisterFactory("default_MinesweeperService", CreateMinesweeperService)

	//bomb locators
	factory.FactoryRegistrySingleton().RegisterFactory("fixed_BombLocator", domain.CreateFixedBombLocator)

	//dals
	factory.FactoryRegistrySingleton().RegisterFactory("inmemory_GameDal", gamedal.CreateInMemoryGameDal)
}


func Test_CreateGameAndLost(t *testing.T) {
	servicesTestSetup()

	service, err := NewMinesweeperService("root.services.minesweeper.board5x5x5")
	helpers.AssertErrorWithMsg(t, "Error creating Minesweeper Service: ", err)

	game, err := service.NewGame("user1", 5, 5, 5)

	helpers.AssertErrorWithMsg(t,"Error creating New Game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Created, "Expected game in Created Status")
	helpers.AssertTrue(t, len(game.Id) > 0, "Expected Id not empty")

	//Its an empty cell
	game, err = service.RevealCell("user1", game.Id, 0,2)
	helpers.AssertErrorWithMsg(t,"Error revealing cell 0, 2", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//Get Game
	game, err = service.GetGame("user1", game.Id)
	helpers.AssertErrorWithMsg(t,"Error getting game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//Its a bomb
	game, err = service.RevealCell("user1", game.Id, 0,0)
	helpers.AssertErrorWithMsg(t,"Error revealing cell 0, 0", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Lost, "Expected game in Lost Status")

	//Get Game
	game, err = service.GetGame("user1", game.Id)
	helpers.AssertErrorWithMsg(t,"Error getting game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Lost, "Expected game in Lost Status")
}

func Test_CreateGameAndWin(t *testing.T) {
	servicesTestSetup()

	service, err := NewMinesweeperService("root.services.minesweeper.board5x5x5")
	helpers.AssertErrorWithMsg(t, "Error creating Minesweeper Service: ", err)

	game, err := service.NewGame("user1", 5, 5, 5)

	helpers.AssertErrorWithMsg(t,"Error creating New Game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Created, "Expected game in Created Status")
	helpers.AssertTrue(t, len(game.Id) > 0, "Expected Id not empty")

	//Its an empty cell
	game, err = service.RevealCell("user1", game.Id, 0,2)
	helpers.AssertErrorWithMsg(t,"Error revealing cell 0, 2", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//Its an empty cell
	game, err = service.RevealCell("user1", game.Id, 2,0)
	helpers.AssertErrorWithMsg(t,"Error revealing cell 2, 0", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//Its an empty cell
	game, err = service.RevealCell("user1", game.Id, 2,4)
	helpers.AssertErrorWithMsg(t,"Error revealing cell 2, 4", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//Its an empty cell
	game, err = service.RevealCell("user1", game.Id, 4,2)
	helpers.AssertErrorWithMsg(t,"Error revealing cell 4, 2", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Won, "Expected game in Won Status")

	//Get Game
	game, err = service.GetGame("user1", game.Id)
	helpers.AssertErrorWithMsg(t,"Error getting game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Won, "Expected game in Won Status")
}

func Test_CreateGameMarkBombsAndWin1(t *testing.T) {
	servicesTestSetup()

	service, err := NewMinesweeperService("root.services.minesweeper.board5x5x5")
	helpers.AssertErrorWithMsg(t, "Error creating Minesweeper Service: ", err)

	game, err := service.NewGame("user1", 5, 5, 5)

	helpers.AssertErrorWithMsg(t,"Error creating New Game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Created, "Expected game in Created Status")
	helpers.AssertTrue(t, len(game.Id) > 0, "Expected Id not empty")

	//Its an empty cell
	game, err = service.MarkCell("user1", game.Id, 0,0, shared.CellMarkType_Flag)
	game, err = service.MarkCell("user1", game.Id, 0,4, shared.CellMarkType_Flag)
	game, err = service.MarkCell("user1", game.Id, 4,0, shared.CellMarkType_Flag)
	game, err = service.MarkCell("user1", game.Id, 4,4, shared.CellMarkType_Flag)
	game, err = service.MarkCell("user1", game.Id, 2,2, shared.CellMarkType_Flag)


	//Reveal all
	for row:=0;row<5;row++ {
		for col:=0;col<5;col++ {
			game, err = service.RevealCell("user1", game.Id, row, col)
			if game.Status == shared.GameStatus_Won {
				break
			}

			helpers.AssertErrorWithMsg(t, "Error revealing cell", err)
		}
	}

	helpers.AssertTrue(t, game.Status == shared.GameStatus_Won, "Expected game in Won Status")

	//Get Game
	game, err = service.GetGame("user1", game.Id)
	helpers.AssertErrorWithMsg(t,"Error getting game", err)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Won, "Expected game in Won Status")
}



