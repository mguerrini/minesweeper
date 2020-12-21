package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/common/apierrors"
	"github.com/minesweeper/src/services"
	"github.com/minesweeper/src/shared"
	"net/http"
)

type MinesweeperController struct {

}

func NewMinesweeperController() *MinesweeperController {
	return &MinesweeperController{}
}

//GET minesweeper/users/:user_id/games/:game_id/show
//for debug
func (this *MinesweeperController)ShowGame(c *gin.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	gameId := c.Param("game_id")
	if gameId == "" {
		return apierrors.NewBadRequest(nil, "Game id is mandatory")
	}

	game, err := services.Singleton().ShowGame(userId, gameId)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)
	c.Set("response", game)

	return nil
}


//GET minesweeper/users/:user_id/games/:game_id
func (this *MinesweeperController)	GetGame(c *gin.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	gameId := c.Param("game_id")
	if gameId == "" {
		return apierrors.NewBadRequest(nil, "Game id is mandatory")
	}

	game, err := services.Singleton().GetGame(userId, gameId)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)
	c.Set("response", game)

	return nil
}


//GET minesweeper/users/:user_id/games
func (this *MinesweeperController)	GetGameListByUserId(c *gin.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	games, err := services.Singleton().GetGameListByUserId(userId)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)
	c.Set("response", games)

	return nil
}


//POST minesweeper/users/:user_id/games
func (this *MinesweeperController)	CreateNewGame(c *gin.Context) error {
	req := &NewGameRequest{}
	err := c.BindJSON(req)
	if err != nil {
		return apierrors.NewBadRequest(err, "Invalid payload from new game")
	}

	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	game, err := services.Singleton().NewGame(userId, req.Rows, req.Columns, req.Mines)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusCreated)
	c.Set("response", game)

	return nil
}


//PUT minesweeper/users/:user_id/games/:game_id/reveal
func (this *MinesweeperController)	RevealCell(c *gin.Context) error {
	req := &RevealCellRequest{}
	err := c.BindJSON(req)
	if err != nil {
		return apierrors.NewBadRequest(err, "Invalid payload from new game")
	}

	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	gameId := c.Param("game_id")
	if gameId == "" {
		return apierrors.NewBadRequest(nil, "Game id is mandatory")
	}

	game, err := services.Singleton().RevealCell(userId, gameId, req.Row, req.Col)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)
	c.Set("response", game)

	return nil
}

//PUT minesweeper/users/:user_id/games/:game_id/mark
func (this *MinesweeperController)	MarkCell(c *gin.Context) error {
	req := &MarkCellRequest{}
	err := c.BindJSON(req)
	if err != nil {
		return apierrors.NewBadRequest(err, "Invalid payload from new game")
	}

	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	gameId := c.Param("game_id")
	if gameId == "" {
		return apierrors.NewBadRequest(nil, "Game id is mandatory")
	}

	if !req.Flag && !req.None && !req.Question {
		return apierrors.NewBadRequest(nil, "At least yout must select one mark (Flag, Question, None)")
	}

	var game *shared.GameData
	if req.Flag {
		game, err = services.Singleton().MarkCell(userId, gameId, req.Row, req.Col, shared.CellMarkType_Flag)
	} else if req.Question {
		game, err = services.Singleton().MarkCell(userId, gameId, req.Row, req.Col, shared.CellMarkType_Question)
	} else {
		game, err = services.Singleton().MarkCell(userId, gameId, req.Row, req.Col, shared.CellMarkType_None)
	}

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)
	c.Set("response", game)

	return nil
}

//DELETE minesweeper/users/:user_id/games/:game_id
func (this *MinesweeperController)	DeleteGame(c *gin.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	gameId := c.Param("game_id")
	if gameId == "" {
		return apierrors.NewBadRequest(nil, "Game id is mandatory")
	}

	err := services.Singleton().DeleteGame(userId, gameId)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)

	return nil
}

//DELETE minesweeper/users/:user_id/games
func (this *MinesweeperController)	DeleteGamesByUser(c *gin.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return apierrors.NewBadRequest(nil, "User id is mandatory")
	}

	err := services.Singleton().DeleteAllGames(userId)

	if err != nil {
		return err
	}

	c.Set("status_code", http.StatusOK)

	return nil
}