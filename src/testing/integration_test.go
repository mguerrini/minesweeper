package testing

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/common/factory"
	"github.com/minesweeper/src/common/helpers"
	"github.com/minesweeper/src/controllers"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
	"github.com/minesweeper/src/server"
	"github.com/minesweeper/src/services"
	"github.com/minesweeper/src/shared"
	"net/http"
	"strconv"
	"testing"
)

func servicesTestSetup() {

	server.StartUp("testing/service_test_configuration.yml")

	//service
	factory.FactoryRegistrySingleton().RegisterFactory("default_MinesweeperService", services.CreateMinesweeperService)

	//mines locators
	factory.FactoryRegistrySingleton().RegisterFactory("fixed_MinesLocator", domain.CreateFixedMinesLocator)

	//dals
	factory.FactoryRegistrySingleton().RegisterFactory("inmemory_GameDal", gamedal.CreateInMemoryGameDal)
}

var UserName string = "cjose"

func Test_CreateGameAndLost(t *testing.T) {
	servicesTestSetup()
	controller := controllers.NewMinesweeperController()

	//engine := server.New()
	game, _ := callNewGameContext(t, controller, 5, 5, 3)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Created, "Expected game in Created Status")

	game, _ = callRevealContext(t, controller, game,0, 2)
	game, _ = callGetGameContext(t, controller, game)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Playing, "Expected game in PLaying Status")

	game, _ = callRevealContext(t, controller, game,2, 2)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Lost, "Expected game in Lost Status")
	game, _ = callGetGameContext(t, controller, game)
	helpers.AssertTrue(t, game.Status == shared.GameStatus_Lost, "Expected game in Lost Status")
}



func callNewGameContext(t *testing.T, controller *controllers.MinesweeperController, rows, cols, mines int) (*shared.GameData, error) {
	ctx := createNewGameContext(rows, cols, mines)
	err := controller.CreateNewGame(ctx)
	helpers.AssertErrorWithMsg(t,"Error creating New Game", err)
	g, _ := ctx.Get("response")
	helpers.AssertTrue(t, g != nil, "Reques not exist in context")

	game := g.(*shared.GameData)
	helpers.AssertErrorWithMsg(t,"Error creating New Game", err)
	helpers.AssertTrue(t, len(game.Id) > 0, "Expected Id not empty")

	return game, err
}

func callRevealContext(t *testing.T, controller *controllers.MinesweeperController, game *shared.GameData, row, col int) (*shared.GameData, error) {
	ctx := createRevealContext(game.Id, row,col)
	err := controller.RevealCell(ctx)
	helpers.AssertErrorWithMsg(t,"Error revealing cell " + strconv.Itoa(row) + ", " + strconv.Itoa(col), err)
	g, _ := ctx.Get("response")
	game = g.(*shared.GameData)

	return game, err
}

func callGetGameContext(t *testing.T, controller *controllers.MinesweeperController, game *shared.GameData) (*shared.GameData, error)  {
	ctx := createGetGameContext(game.Id)
	err := controller.GetGame(ctx)
	helpers.AssertErrorWithMsg(t, "Error getting game", err)
	g, _ := ctx.Get("response")
	game = g.(*shared.GameData)

	return game, err
}


func createNewGameContext(rows, cols, mines int) *gin.Context {
	//New Game
	req := controllers.NewGameRequest{
		Rows:    rows,
		Columns: cols,
		Mines:   mines,
	}
	body, _ := json.Marshal(req)

	c := &gin.Context{}

	c.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(body))
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: UserName}}

	return c
}

func createRevealContext(gameId string, row, col int) *gin.Context {
	//New Game
	req := &controllers.RevealCellRequest{
		controllers.BaseGameActionRequest {
			Row: row,
			Col: col,
		},
	}

	body, _ := json.Marshal(req)

	c := &gin.Context{}

	c.Request, _ = http.NewRequest("PUT", "", bytes.NewBuffer(body))
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: UserName}, gin.Param{Key: "game_id", Value: gameId}}

	return c
}

func createGetGameContext(gameId string) *gin.Context {

	c := &gin.Context{}

	c.Request, _ = http.NewRequest("PUT", "", nil)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: UserName}, gin.Param{Key: "game_id", Value: gameId}}

	return c
}
