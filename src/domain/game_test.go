package domain

import (
	"fmt"
	"github.com/minesweeper/src/common/helpers"
	"github.com/minesweeper/src/shared"
	"testing"
	"time"
)

func createGameFactory1() MinesweeperGameFactory {
	mineLoc := NewFixedMineLocator()

	mineLoc.AddMine(0,0)
	mineLoc.AddMine(0,4)
	mineLoc.AddMine(4,0)
	mineLoc.AddMine(4,4)
	mineLoc.AddMine(2,2)

	return NewMinesweeperGameFactory(mineLoc)
}

func createGameFactory2() MinesweeperGameFactory {
	mineLoc := NewFixedMineLocator()

	mineLoc.AddMine(2,2)
	mineLoc.AddMine(4,0)
	mineLoc.AddMine(4,4)

	return NewMinesweeperGameFactory(mineLoc)
}

func Test_CreateGame_GetData(t *testing.T) {
	factory := createGameFactory1()
	game, err := factory.CreateGame(5, 5, 5)

	helpers.AssertError(t, err)

	data := game.GetData()

	helpers.AssertTrue(t, data.StartTime == time.Time{}, "Expected 0 start time")
	helpers.AssertTrue(t, data.FinishTime == time.Time{}, "Expected 0 finish time")
	helpers.AssertTrue(t, data.Status == shared.GameStatus_Created, "Expected game in Created Status")
	helpers.AssertTrue(t, data.Board.RowCount == 5, "Expected 5 rows")
	helpers.AssertTrue(t, data.Board.ColCount == 5, "Expected 5 cols")
	helpers.AssertTrue(t, data.Board.MinesCount == 5, "Expected 5 mines")
	helpers.AssertTrue(t, len(data.Board.Cells) == 5, "Expected 5 cell per row")
	helpers.AssertTrue(t, len(data.Board.Cells[0]) == 5, "Expected 5 cell in col 0")
	helpers.AssertTrue(t, len(data.Board.Cells[1]) == 5, "Expected 5 cell in col 1")
	helpers.AssertTrue(t, len(data.Board.Cells[2]) == 5, "Expected 5 cell in col 2")
	helpers.AssertTrue(t, len(data.Board.Cells[3]) == 5, "Expected 5 cell in col 3")
	helpers.AssertTrue(t, len(data.Board.Cells[4]) == 5, "Expected 5 cell in col 4")

	//check mines
	validateMineCell(t, &data, 0,0, false)
	validateMineCell(t, &data, 0,4, false)
	validateMineCell(t, &data, 4,0, false)
	validateMineCell(t, &data, 4,4, false)
	validateMineCell(t, &data, 2,2, false)

	//check numbers
	validateNumberCell(t, &data, 0,1, 1, false)
	validateNumberCell(t, &data, 1,1, 2, false)
	validateNumberCell(t, &data, 1,0, 1, false)

	validateNumberCell(t, &data, 0,3, 1, false)
	validateNumberCell(t, &data, 1,3, 2, false)
	validateNumberCell(t, &data, 1,4, 1, false)

	validateNumberCell(t, &data, 3,0, 1, false)
	validateNumberCell(t, &data, 3,1, 2, false)
	validateNumberCell(t, &data, 4,1, 1, false)

	validateNumberCell(t, &data, 3,4, 1, false)
	validateNumberCell(t, &data, 3,3, 2, false)
	validateNumberCell(t, &data, 4,3, 1, false)

	validateNumberCell(t, &data, 1,2, 1, false)
	validateNumberCell(t, &data, 2,1, 1, false)
	validateNumberCell(t, &data, 2,3, 1, false)
	validateNumberCell(t, &data, 3,2, 1, false)

	//Empty cells
	validateEmptyCell(t, &data, 2,0,  false)
	validateEmptyCell(t, &data, 2,4,  false)
	validateEmptyCell(t, &data, 0,2,  false)
	validateEmptyCell(t, &data, 4,2,  false)
}

func Test_HideBoardData(t *testing.T){
	factory := createGameFactory1()
	game, _ := factory.CreateGame(5, 5, 5)

	data:=game.GetData()

	data.Board.HideNotRevealed()

	//all cells must be the type unknown without marks
	for r := range data.Board.Cells {
		for c, _ := range data.Board.Cells[r] {
			cellData := data.Board.GetCell(r, c)
			helpers.AssertTrue(t, cellData.Type == shared.CellType_Unknown, fmt.Sprintf("Expected Cell Type == Unknown type for cell %d, %d", r, c))
			helpers.AssertTrue(t, cellData.IsRevealed == false, fmt.Sprintf("Expected IsRevealed == false for cell %d, %d", r, c))
			helpers.AssertTrue(t, cellData.Mark == shared.CellMarkType_None, fmt.Sprintf("Expected Mark == None, for cell %d, %d", r, c))
		}
	}

}

func Test_RevealEmptyCell(t *testing.T){
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	err := game.RevealCell(0, 0)
	helpers.AssertError(t, err)

	//get data and hide not revealed
	data:=game.GetData()

	//check game status
	helpers.AssertTrue(t, data.StartTime != time.Time{}, "Expected start time != 0")
	helpers.AssertTrue(t, data.FinishTime == time.Time{}, "Expected 0 finish time")
	helpers.AssertTrue(t, data.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//hide cells
	data.Board.HideNotRevealed()

	//check cell type
	validateEmptyCell(t, &data, 0, 0, true)
	validateEmptyCell(t, &data, 0, 1, true)
	validateEmptyCell(t, &data, 0, 2, true)
	validateEmptyCell(t, &data, 0, 3, true)
	validateEmptyCell(t, &data, 0, 4, true)

	validateEmptyCell(t, &data, 1,0, true)
	validateEmptyCell(t, &data, 2,0, true)
	validateEmptyCell(t, &data, 1,4, true)
	validateEmptyCell(t, &data, 2,4, true)

	validateNumberCell(t, &data, 1,1, 1, true)
	validateNumberCell(t, &data, 1,2, 1,true)
	validateNumberCell(t, &data, 1,3, 1,true)
	validateNumberCell(t, &data, 2,1, 1,true)
	validateNumberCell(t, &data, 2,3, 1,true)
	validateNumberCell(t, &data, 3,0, 1,true)
	validateNumberCell(t, &data, 3,1, 2,true)
	validateNumberCell(t, &data, 3,3, 2,true)
	validateNumberCell(t, &data, 3,4, 1,true)

	validateUnknownCell(t, &data, 2,2, false)
	validateUnknownCell(t, &data, 3,2, false)
	validateUnknownCell(t, &data, 4,0, false)
	validateUnknownCell(t, &data, 4,1, false)
	validateUnknownCell(t, &data, 4,2, false)
	validateUnknownCell(t, &data, 4,3, false)
	validateUnknownCell(t, &data, 4,4, false)
}

func Test_LostGameCell(t *testing.T){
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	//reveal empty cell
	err := game.RevealCell(0, 0)
	helpers.AssertError(t, err)

	//check game status
	helpers.AssertTrue(t, game.data.StartTime != time.Time{}, "Expected start time != 0")
	helpers.AssertTrue(t, game.data.FinishTime == time.Time{}, "Expected 0 finish time")
	helpers.AssertTrue(t, game.data.Status == shared.GameStatus_Playing, "Expected game in Playing Status")

	//reveal a mine
	err = game.RevealCell(2, 2)
	helpers.AssertError(t, err)

	//check game status
	helpers.AssertTrue(t, game.data.StartTime != time.Time{}, "Expected start time != 0")
	helpers.AssertTrue(t, game.data.FinishTime != time.Time{}, "Expected finish time != 0")
	helpers.AssertTrue(t, game.data.Status == shared.GameStatus_Lost, "Expected game in Lost Status")

	//get data and hide not revealed
	data:=game.GetData()

	//check cell type
	for row := range data.Board.Cells {
		for col := range data.Board.Cells[row] {
			if (row == 2 && col == 2) || (row == 4 && (col == 0 || col == 4)) {
				//mine
				validateMineCell(t, &data, row, col, true)
			} else {
				validateEmptyCell(t, &data, row, col, true)
			}
		}
	}
}

func Test_WinGameCell(t *testing.T){
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	err := game.RevealCell(0, 0)
	helpers.AssertError(t, err)
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	err = game.RevealCell(3, 2)
	helpers.AssertError(t, err)
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	err = game.RevealCell(4, 2)
	helpers.AssertError(t, err)
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Won, "Expected game in Won Status")
	helpers.AssertTrue(t, game.data.StartTime != time.Time{}, "Expected start time != 0")
	helpers.AssertTrue(t, game.data.FinishTime != time.Time{}, "Expected finish time != 0")

	data := game.GetData()

	//check cell type
	for row := range data.Board.Cells {
		for col := range data.Board.Cells[row] {
			if (row == 2 && col == 2) || (row == 4 && (col == 0 || col == 4)) {
				//mine
				validateMineCell(t, &data, row, col, true)
			} else {
				validateEmptyCell(t, &data, row, col, true)
			}
		}
	}
}

func Test_MarkWithFlag(t *testing.T) {
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	//mark a mine with a falg and try to revealed
	err:=game.MarkCell(2,2, shared.CellMarkType_Flag)
	helpers.AssertError(t, err)

	//start game
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	//check cell
	data:=game.GetData()
	data.Board.HideNotRevealed()
	validateUnknownCell(t, &data, 2, 2, false)
	validateFlagMark(t, &data, 2, 2)

	//try to reveal mine...game not change
	err = game.RevealCell(2, 2)
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	data=game.GetData()
	data.Board.HideNotRevealed()
	validateUnknownCell(t, &data, 2, 2, false)
	validateFlagMark(t, &data, 2, 2)

	//unmark cell
	err =game.MarkCell(2,2, shared.CellMarkType_None)
	helpers.AssertError(t, err)

	//check game status
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")
	data=game.GetData()
	data.Board.HideNotRevealed()
	validateUnknownCell(t, &data, 2, 2, false)
	validateNoneMark(t, &data, 2, 2)

	//reveal a mine
	err = game.RevealCell(2, 2)
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Lost, "Expected game in Lost Status")
}

func Test_MarkWithQuestionAMine(t *testing.T) {
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	//mark a mine with a falg and try to revealed
	err:=game.MarkCell(2,2, shared.CellMarkType_Question)
	helpers.AssertError(t, err)

	//start game
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	//check cell
	data:=game.GetData()
	data.Board.HideNotRevealed()
	validateUnknownCell(t, &data, 2, 2, false)
	validateQuestionMark(t, &data, 2, 2)

	//reveal a mine
	err = game.RevealCell(2, 2)
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Lost, "Expected game in Lost Status")
}

func Test_MarkWithQuestionAnEmptyCell(t *testing.T) {
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	//mark a mine with a falg and try to revealed
	err:=game.MarkCell(0,0, shared.CellMarkType_Question)
	helpers.AssertError(t, err)

	//start game
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	//check cell
	data:=game.GetData()
	data.Board.HideNotRevealed()
	validateUnknownCell(t, &data, 0, 0, false)
	validateQuestionMark(t, &data, 0, 0)

	//reveal a mine
	err = game.RevealCell(0, 0)
	data =game.GetData()
	data.Board.HideNotRevealed()

	validateEmptyCell(t, &data, 0, 0, true)
	validateNoneMark(t, &data, 0, 0)
}

func Test_MarkWithQuestionAnNumberCell(t *testing.T) {
	factory := createGameFactory2()
	game, _ := factory.CreateGame(5, 5, 3)

	//mark a mine with a falg and try to revealed
	err:=game.MarkCell(1,1, shared.CellMarkType_Question)
	helpers.AssertError(t, err)

	//start game
	helpers.AssertTrue(t, game.GetStatus() == shared.GameStatus_Playing, "Expected game in Playing Status")

	//check cell
	data:=game.GetData()
	data.Board.HideNotRevealed()
	validateUnknownCell(t, &data, 1, 1, false)
	validateQuestionMark(t, &data, 1, 1)

	//reveal a number
	err = game.RevealCell(1, 1)
	data =game.GetData()
	data.Board.HideNotRevealed()

	validateNumberCell(t, &data, 1, 1, 1,true)
	validateNoneMark(t, &data, 1, 1)
}

func validateEmptyCell(t *testing.T, game *shared.GameData, row, col int, isRevealedExpectedValue bool) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.IsRevealed == isRevealedExpectedValue, fmt.Sprintf("Expected IsRevealed == %t for cell %d, %d", isRevealedExpectedValue, row, col))
	helpers.AssertTrue(t, cellData.Type == shared.CellType_Empty, fmt.Sprintf("Expected Cell Type == Empty type for cell %d, %d", row, col))
}

func validateNumberCell(t *testing.T, game *shared.GameData, row, col int, number int, isRevealedExpectedValue bool) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.IsRevealed == isRevealedExpectedValue, fmt.Sprintf("Expected IsRevealed == %t for cell %d, %d", isRevealedExpectedValue, row, col))
	helpers.AssertTrue(t, cellData.Type == shared.CellType_Number, fmt.Sprintf("Expected Cell Type == Number type for cell %d, %d", row, col))
	helpers.AssertTrue(t, cellData.Number == number, fmt.Sprintf("Expected Number %d in cell %d, %d", number, row, col))
}

func validateUnknownCell(t *testing.T, game *shared.GameData, row, col int, isRevealedExpectedValue bool) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.IsRevealed == isRevealedExpectedValue, fmt.Sprintf("Expected IsRevealed == %t for cell %d, %d", isRevealedExpectedValue, row, col))
	helpers.AssertTrue(t, cellData.Type == shared.CellType_Unknown, fmt.Sprintf("Expected Cell Type == Unknown type for cell %d, %d", row, col))
}

func validateMineCell(t *testing.T, game *shared.GameData, row, col int, isRevealedExpectedValue bool) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.IsRevealed == isRevealedExpectedValue, fmt.Sprintf("Expected IsRevealed == %t for cell %d, %d", isRevealedExpectedValue, row, col))
	helpers.AssertTrue(t, cellData.Type == shared.CellType_Mine, fmt.Sprintf("Expected Cell Type == Mine type for cell %d, %d", row, col))
}

func validateFlagMark(t *testing.T, game *shared.GameData, row, col int) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.Type == shared.CellType_Unknown, fmt.Sprintf("Expected Cell Type == Unknown type for cell %d, %d", row, col))
	helpers.AssertTrue(t, cellData.Mark == shared.CellMarkType_Flag, fmt.Sprintf("Expected Cell Mark == Flag type for cell %d, %d", row, col))
}

func validateQuestionMark(t *testing.T, game *shared.GameData, row, col int) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.Type == shared.CellType_Unknown, fmt.Sprintf("Expected Cell Type == Unknown type for cell %d, %d", row, col))
	helpers.AssertTrue(t, cellData.Mark == shared.CellMarkType_Question, fmt.Sprintf("Expected Cell Mark == Question type for cell %d, %d", row, col))
}

func validateNoneMark(t *testing.T, game *shared.GameData, row, col int) {
	cellData := game.Board.GetCell(row, col)
	helpers.AssertTrue(t, cellData.Mark == shared.CellMarkType_None, fmt.Sprintf("Expected Cell Mark == None type for cell %d, %d", row, col))
}


