package domain

import (
	"github.com/minesweeper/src/shared"
)

type Cell interface {
	GetData() shared.CellData
	GetType() shared.CellType
	IsRevealed() bool
	IsMarked() bool

	setMark(markType shared.CellMarkType)
	setRevealed(isRevealed bool)
	Mark(board *Board, mark shared.CellMarkType)
	Reveal(board *Board) bool
}


/***********************************/
/*         Base Cell               */
/***********************************/

type cell struct {
	data shared.CellData
}


func NewEmptyCell(row, col int) Cell{
	return &cell{
		data:  shared.CellData{
			Type:       shared.CellType_Empty,
			Row:        row,
			Col:        col,
			IsRevealed: false,
			Mark:       shared.CellMarkType_None,
			Number:     -1,
		},
	}
}

func (this *cell) GetData () shared.CellData {
	return this.data
}

func (this *cell) GetType () shared.CellType {
	return this.data.Type
}

func (this *cell) IsRevealed () bool {
	return this.data.IsRevealed
}

func (this *cell) IsMarked () bool {
	return false
}

func (this *cell) Reveal(board *Board) bool  {
	if this.IsRevealed() {
		return false
	}

	this.data.IsRevealed = true;

	//empty - expose the neighbors
	for row:= this.data.Row-1; row<= this.data.Row+1; row++ {
		for col:= this.data.Col-1; col<= this.data.Col+1; col++ {
			//check limits
			if row >= 0 && row < board.GetMaxRow() && col >= 0 && col < board.GetMaxCol(){
				nCell := board.getCell(row, col)

				if nCell.GetType() != shared.CellType_Mine {
					nCell.Reveal(board)
				}
			}
		}
	}

	return false
}

func (this *cell) Mark(board *Board, mark shared.CellMarkType)  {
	this.doMark(board, this, mark)
}

func (this *cell) doMark(board *Board, toMark Cell, mark shared.CellMarkType)  {
	if toMark.IsRevealed() {
		return
	}

	if mark == shared.CellMarkType_None {
		return
	}

	//change cell
	mCell := NewMarkedCell(toMark)
	innerData := toMark.GetData()
	board.setCell(innerData.Row, innerData.Col, mCell)

	mCell.Mark(board, mark)
}


func (this *cell) setMark(mark shared.CellMarkType)   {
	this.data.Mark = mark
}

func (this *cell) setRevealed(isRevealed bool)   {
	this.data.IsRevealed = isRevealed
}


/***********************************/
/*         Marked Cell               */
/***********************************/

type markedCell struct {
	cell
	markedCell Cell
}

func NewMarkedCell(innerCell Cell) Cell {
	return &markedCell{
		markedCell: innerCell,
	}
}


func (this *markedCell) GetData () shared.CellData {
	return this.markedCell.GetData()
}

func (this *markedCell) GetType () shared.CellType {
	return this.markedCell.GetType()
}

func (this *markedCell) IsRevealed () bool {
	return false
}

func (this *markedCell) IsMarked () bool {
	return true
}

func (this *markedCell) Reveal(board *Board) bool  {
	innerData := this.markedCell.GetData()

	if innerData.Mark == shared.CellMarkType_Flag {
		return false
	}

	//change de cell in the board.....put the original cell and clean the mark
	this.markedCell.setMark(shared.CellMarkType_None)
	board.setCell(innerData.Row, innerData.Col, this.markedCell)

	//then reveal de new cell
	output := this.markedCell.Reveal(board)

	return output
}

func (this *markedCell) Mark(board *Board, mark shared.CellMarkType)  {
	//put the mark in the cell
	this.markedCell.setMark(mark)

	if mark == shared.CellMarkType_None {
		innerData := this.markedCell.GetData()
		board.setCell(innerData.Row, innerData.Col, this.markedCell)
	}
}

func (this *markedCell) setMark(mark shared.CellMarkType)   {
	this.markedCell.setMark(mark)
}

func (this *markedCell) setRevealed(isRevealed bool)   {
	this.markedCell.setRevealed(isRevealed)
}


/***********************************/
/*         Mine Cell               */
/***********************************/

type mineCell struct {
	cell
}

func NewMineCell(row, col int) Cell {
	return &mineCell{
		cell: cell{
			data: shared.CellData{
				Type:       shared.CellType_Mine,
				Row:        row,
				Col:        col,
				IsRevealed: false,
				Mark:       shared.CellMarkType_None,
				Number:     -1,
			},
		},
	}
}

func (this *mineCell) Reveal(board *Board) bool {
	if this.IsRevealed() {
		return false
	}

	this.data.IsRevealed = true;

	return true
}

func (this *mineCell) Mark(board *Board, mark shared.CellMarkType)  {
	this.doMark(board, this, mark)
}


/***********************************/
/*         Number Cell               */
/***********************************/

type numberCell struct {
	cell
}

func NewNumberCell(row, col int, number int) Cell{
	return &numberCell{
		cell: cell{
			data: shared.CellData{
				Type:       shared.CellType_Number,
				Row:        row,
				Col:        col,
				IsRevealed: false,
				Mark:       shared.CellMarkType_None,
				Number:     number,
			},
		},
	}
}

func (this *numberCell) Reveal(board *Board) bool {
	if this.IsRevealed() {
		return false
	}

	this.data.IsRevealed = true;

	return false
}

func (this *numberCell) Mark(board *Board, mark shared.CellMarkType)  {
	this.doMark(board, this, mark)
}



