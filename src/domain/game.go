package domain

import (
	"errors"
	"fmt"
	"github.com/minesweeper/src/shared"
	"time"
)

type Game struct {
	data shared.GameData
	board Board
}


func NewGame(rowCount, colCount int, bombCount int, bombLocator BombLocator) Game {
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

	bombLocator.SetBombs(&game, bombCount)

	game.board.setCellNumbers()

	return game
}


func (this *Game) GetId() string {
	return this.data.Id
}

func (this *Game) SetId(id string)  {
	this.data.Id = id
}


func (this *Game) GetData() shared.GameData {
	copy := this.data

	//complete fields
	copy.Board = this.board.getData();

	return copy
}

func (this *Game) GetRowCount() int {
	return this.board.GetMaxRow()
}

func (this *Game) GetColCount() int {
	return this.board.GetMaxCol()
}


func (this *Game) IsFinished() bool  {
	return this.data.Status == shared.GameStatus_Lost || this.data.Status == shared.GameStatus_Won
}

func (this *Game) SetBomb(row int, col int) (bool, error) {
	if this.data.Status == shared.GameStatus_Created {
		return this.board.SetBomb(row, col), nil
	}

	return false, errors.New("The game is started, can not add more bombs")
}

func (this *Game) RevealCell(row int, col int) error {
	err := this.areInRange(row, col)

	if err != nil {
		return err
	}

	//if is the first cell exposed => start clock
	count := this.board.GetRevealedCount()

	if count == 0 {
		this.data.StartTime = time.Now()
		this.data.Status = shared.GameStatus_Playing
	}

	if this.data.Status != shared.GameStatus_Playing {
		return errors.New("The game is not started!")
	}

	cell := this.board.getCell(row, col)

	isBomb := cell.Expose(&this.board)

	if isBomb	{
		//game end
		this.data.FinishTime = time.Now()
		this.data.Status = shared.GameStatus_Lost
	}

	//check if won
	notExposedCount := this.board.GetNotRevealedCount()

	if notExposedCount == this.board.GetBombsCount() {
		//won the game
		this.data.FinishTime = time.Now()
		this.data.Status = shared.GameStatus_Won
	}

	return nil
}

func (this *Game) MarkCell(row int, col int) error {
	err := this.areInRange(row, col)

	if err != nil {
		return err
	}

	//if is the first cell exposed => start clock
	count := this.board.GetRevealedCount()

	if count == 0 {
		this.data.StartTime = time.Now()
		this.data.Status = shared.GameStatus_Playing
	}

	if this.data.Status != shared.GameStatus_Playing {
		return errors.New("The game is not started!")
	}

	cell := this.board.getCell(row, col)

	cell.Mark()

	return nil
}

func (this *Game) areInRange(row int, col int) error {
	if row < 0 || row >= this.board.data.RowCount {
		return errors.New(fmt.Sprintf("Invalid row number. Its must be between 0 and %d", this.board.data.RowCount))
	}

	if col < 0 || col >= this.board.data.ColCount {
		return errors.New(fmt.Sprintf("Invalid col number. Its must be between 0 and %d", this.board.data.ColCount))
	}

	return nil
}


