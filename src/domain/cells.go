package domain

import (
	"github.com/minesweeper/src/shared"
)

type Cell interface {
	GetData() shared.CellData
	GetType() shared.CellType
	IsExposed() bool
	IsMarked() bool

	Mark()
	Expose(board *Board) bool
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
			Type:      shared.CellType_Empty,
			Row:       row,
			Col:       col,
			IsReveled: false,
			IsMarked:  false,
			Number:    -1,
		},
	}
}

func (this *cell) GetData () shared.CellData {
	return this.data
}

func (this *cell) GetType () shared.CellType {
	return this.data.Type
}

func (this *cell) IsExposed () bool {
	return this.data.IsReveled
}

func (this *cell) IsMarked () bool {
	return this.data.IsMarked
}

func (this *cell) Expose(board *Board) bool  {
	if this.IsExposed() || this.IsMarked(){
		return false
	}

	this.data.IsReveled = true;

	//empty - expose the neighbors
	for row:= this.data.Row-1; row<= this.data.Row+1; row++ {
		for col:= this.data.Col-1; col<= this.data.Col+1; col++ {
			//check limits
			if row >= 0 && row < board.GetMaxRow() && col >= 0 && col < board.GetMaxCol(){
				nCell := board.getCell(row, col)

				if nCell.GetType() != shared.CellType_Bomb {
					nCell.Expose(board)
				}
			}
		}
	}

	return false
}

func (this *cell) Mark()  {
	if this.IsExposed() {
		return
	}

	this.data.IsMarked = !this.data.IsMarked
}

/***********************************/
/*         Bomb Cell               */
/***********************************/

type bombCell struct {
	cell
}

func NewBombCell(row, col int) Cell {
	return &bombCell{
		cell: cell{
			data: shared.CellData{
				Type:      shared.CellType_Bomb,
				Row:       row,
				Col:       col,
				IsReveled: false,
				IsMarked:  false,
				Number:    -1,
			},
		},
	}
}

func (this *bombCell) Expose(board *Board) bool {
	if this.IsExposed() || this.IsMarked(){
		return false
	}

	this.data.IsReveled = true;

	return true
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
				Type:      shared.CellType_Number,
				Row:       row,
				Col:       col,
				IsReveled: false,
				IsMarked:  false,
				Number:    number,
			},
		},
	}
}

func (this *numberCell) Expose(board *Board) bool {
	if this.IsExposed() || this.IsMarked(){
		return false
	}

	this.data.IsReveled = true;

	return false
}




