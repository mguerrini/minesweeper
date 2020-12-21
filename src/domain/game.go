package domain

import (
	"fmt"
	"github.com/minesweeper/src/common/apierrors"
	"github.com/minesweeper/src/shared"
	"time"
)

type Game struct {
	data shared.GameData
	board Board
}


func NewGame(rowCount, colCount int, bombCount int, bombLocator BombLocator) (Game, error) {
	board := NewBoard(rowCount, colCount)

	game := Game{
		data:  shared.GameData{
			Id:         "",
			StartTime:  time.Time{},
			FinishTime: time.Time{},
			Status:     shared.GameStatus_Created,
		},
		board: board,
	}

	err := bombLocator.SetBombs(&game, bombCount)

	if err != nil {
		return Game{}, err
	}

	game.board.setCellNumbers()

	return game, nil
}


func (this *Game) GetId() string {
	return this.data.Id
}

func (this *Game) SetId(id string)  {
	this.data.Id = id
}

func (this *Game) GetRowCount() int {
	return this.data.Board.RowCount
}

func (this *Game) GetStatus() shared.GameStatusType {
	return this.data.Status
}

func (this *Game) GetColCount() int  {
	return this.data.Board.ColCount
}


func (this *Game) GetData() shared.GameData {
	copy := this.data

	//complete fields
	copy.Board = this.board.getData();

	return copy
}


func (this *Game) IsFinished() bool  {
	return this.data.Status == shared.GameStatus_Lost || this.data.Status == shared.GameStatus_Won
}

func (this *Game) SetBomb(row int, col int) (bool, error) {
	if err :=this.areInRange(row, col); err != nil {
		return false, err
	}

	if this.data.Status == shared.GameStatus_Created {
		return this.board.SetBomb(row, col), nil
	}

	return false, apierrors.NewBadRequest(nil, "The game is started, can not add more bombs")
}

func (this *Game) RevealCell(row int, col int) error {
	if err := this.areInRange(row, col); err != nil {
		return err
	}

	err := this.startGame()
	if err != nil {
		return err
	}

	cell := this.board.getCell(row, col)

	isBomb := cell.Reveal(&this.board)

	if isBomb	{
		//game end
		this.gameOver(false)
	}

	//check if won
	notExposedCount := this.board.GetNotRevealedCount()

	if notExposedCount == this.board.GetBombsCount() {
		this.gameOver(true)
	}

	return nil
}

func (this *Game) MarkCell(row int, col int, mark shared.CellMarkType) error {
	if err := this.areInRange(row, col); err != nil {
		return err
	}

	err := this.startGame()
	if err != nil {
		return err
	}

	cell := this.board.getCell(row, col)

	cell.Mark(&this.board, mark)

	return nil
}




func (this *Game) startGame () error  {

	if this.data.Status == shared.GameStatus_Playing {
		return nil
	} else if this.data.Status == shared.GameStatus_Created {
		this.data.StartTime = time.Now()
		this.data.Status = shared.GameStatus_Playing
		return nil
	}

	return apierrors.NewBadRequest(nil, "The game is finished!")
}

func (this *Game) gameOver (won bool) {
	//revelead all cell and leave only de bombs
	this.data.FinishTime = time.Now()

	if won {
		//won the game
		this.data.Status = shared.GameStatus_Won
	} else {
		this.data.Status = shared.GameStatus_Lost
	}

	this.board.revealBombs()
}

func (this *Game) areInRange(row int, col int) error {
	if row < 0 || row >= this.board.data.RowCount {
		return apierrors.NewBadRequest(nil, fmt.Sprintf("Invalid row number. Its must be between 0 and %d", this.board.data.RowCount))
	}

	if col < 0 || col >= this.board.data.ColCount {
		return apierrors.NewBadRequest(nil, fmt.Sprintf("Invalid col number. Its must be between 0 and %d", this.board.data.ColCount))
	}

	return nil
}

