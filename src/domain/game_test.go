package domain

import (
	"github.com/minesweeper/src/common/helpers"
	"github.com/minesweeper/src/shared"
	"testing"
	"time"
)

func createGameFactory() MinesweeperGameFactory {
	bombLoc := NewFixedBombLocator()

	bombLoc.AddBomb(0,0)
	bombLoc.AddBomb(0,4)
	bombLoc.AddBomb(4,0)
	bombLoc.AddBomb(4,4)
	bombLoc.AddBomb(2,2)

	return NewMinesweeperGameFactory(bombLoc)
}

func Test_CreateGame_GetData(t *testing.T) {
	factory := createGameFactory()
	game, err := factory.CreateGame(5, 5, 5)

	helpers.AssertError(t, err)

	data := game.GetData()

	helpers.AssertTrue(t, data.StartTime == time.Time{}, "Expected 0 start time")
	helpers.AssertTrue(t, data.FinishTime == time.Time{}, "Expected 0 finish time")
	helpers.AssertTrue(t, data.Status == shared.GameStatus_Created, "Expected game in Created Status")
	helpers.AssertTrue(t, data.Board.RowCount == 5, "Expected 5 rows")
	helpers.AssertTrue(t, data.Board.ColCount == 5, "Expected 5 cols")
	helpers.AssertTrue(t, data.Board.BombsCount == 5, "Expected 5 bombs")
	helpers.AssertTrue(t, len(data.Board.Cells) == 5, "Expected 5 cell per row")
	helpers.AssertTrue(t, len(data.Board.Cells[0]) == 5, "Expected 5 cell in col 0")
	helpers.AssertTrue(t, len(data.Board.Cells[1]) == 5, "Expected 5 cell in col 1")
	helpers.AssertTrue(t, len(data.Board.Cells[2]) == 5, "Expected 5 cell in col 2")
	helpers.AssertTrue(t, len(data.Board.Cells[3]) == 5, "Expected 5 cell in col 3")
	helpers.AssertTrue(t, len(data.Board.Cells[4]) == 5, "Expected 5 cell in col 4")

	//check bombs
	cell := data.Board.GetCell(0, 0)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Bomb, "Expected a bomb in cell 0,0")
	cell = data.Board.GetCell(0, 4)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Bomb, "Expected a bomb in cell 0,4")
	cell = data.Board.GetCell(4, 0)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Bomb, "Expected a bomb in cell 4,0")
	cell = data.Board.GetCell(4, 4)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Bomb, "Expected a bomb in cell 4,4")
	cell = data.Board.GetCell(2, 2)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Bomb, "Expected a bomb in cell 2,2")

	//check numbers
	cell = data.Board.GetCell(0, 1)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 0,1")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 0,1")

	cell = data.Board.GetCell(1, 1)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 1,1")
	helpers.AssertTrue(t, cell.Number == 2, "Expected Number 2 in cell 1,1")

	cell = data.Board.GetCell(1, 0)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 1,0")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 1,0")


	cell = data.Board.GetCell(0, 3)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 0,3")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 0,3")

	cell = data.Board.GetCell(1, 3)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 1,3")
	helpers.AssertTrue(t, cell.Number == 2, "Expected Number 2 in cell 1,3")

	cell = data.Board.GetCell(1, 4)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 1,4")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 0,4")


	cell = data.Board.GetCell(3, 0)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 3,0")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 3,0")

	cell = data.Board.GetCell(3, 1)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 3,1")
	helpers.AssertTrue(t, cell.Number == 2, "Expected Number 2 in cell 3,1")

	cell = data.Board.GetCell(4, 1)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 4,1")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 4,1")


	cell = data.Board.GetCell(3, 3)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 3,3")
	helpers.AssertTrue(t, cell.Number == 2, "Expected Number 2 in cell 3,3")

	cell = data.Board.GetCell(3, 4)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 3,4")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 3,4")

	cell = data.Board.GetCell(4, 3)
	helpers.AssertTrue(t, cell.Type == shared.CellType_Number, "Expected a Number in cell 4,3")
	helpers.AssertTrue(t, cell.Number == 1, "Expected Number 1 in cell 4,3")
}
