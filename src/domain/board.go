package domain

import "github.com/minesweeper/src/shared"

type Board struct {
	data shared.BoardData
	cells [][]Cell
}

func NewBoard(rowCount, colCount int) Board {
	cells := make([][]Cell, rowCount)
	for row := range cells {
		cells[row] = make([]Cell, colCount)
	}

	cellDatas := make([][]shared.CellData, rowCount)
	for row := range cellDatas {
		cellDatas[row] = make([]shared.CellData, colCount)
	}

	board := Board{
		data:  shared.BoardData{
			RowCount:   rowCount,
			ColCount:   colCount,
			Cells:      cellDatas,
			BombsCount: 0,
		},
		cells: cells,
	}

	board.initializeCells()

	return board
}

func (this *Board) getData() shared.BoardData {
	copy := this.data

	//I go through the cells to refresh copy
	for row := 0; row<this.data.RowCount; row++{
		for col := 0; col<this.data.ColCount; col++{
			cell := this.getCell(row, col)
			copy.SetCell(row, col, cell.GetData())
		}
	}

	return copy
}

func (this *Board) initializeCells() {
	for row := range this.cells {
		for col, _ := range this.cells[row] {
			nCell := NewEmptyCell(row, col)
			this.setCell(row, col, nCell)
		}
	}
}

func (this *Board) GetMaxRow() int{
	return this.data.RowCount
}

func (this *Board) GetMaxCol() int{
	return this.data.ColCount
}

func (this *Board) GetBombsCount() int{
	return this.data.BombsCount
}

func (this *Board) GetRevealedCount() int{
	output := 0

	for row := 0; row<this.data.RowCount; row++{
		for col := 0; row<this.data.ColCount; col++{
			cell := this.getCell(row, col)
			if cell.IsRevealed() {
				output++
			}
		}
	}

	return output
}

func (this *Board) GetNotRevealedCount() int{
	output := 0

	for row := 0; row<this.data.RowCount; row++{
		for col := 0; row<this.data.ColCount; col++{
			cell := this.getCell(row, col)
			if !cell.IsRevealed() {
				output++
			}
		}
	}

	return output
}



func (this *Board) SetBomb(row int, col int) bool {
	cell := this.getCell(row, col)

	if cell.GetType() != shared.CellType_Bomb {
		//convert this cell to bomb cell
		this.setCell(row, col, NewBombCell(row, col))
		this.data.BombsCount++
		return true
	}

	return false
}

func (this *Board) getCell(row int, col int) Cell {
	return this.cells[row][col]
}

func (this *Board) setCell(row int, col int, cell Cell)  {
	this.cells[row][col] = cell
}

func (this *Board) setCellNumbers()  {
	for row := 0; row<this.data.RowCount; row++ {
		for col := 0; col<this.data.ColCount; col++ {
			cell:= this.getCell(row, col)

			if cell.GetType() == shared.CellType_Empty {
				count := this.getCountNeighboringBombs(cell)

				if count > 0 {
					numberCell:= NewNumberCell(row, col, count)
					this.setCell(row, col, numberCell)
				}
			}
		}
	}
}

func (this *Board) getCountNeighboringBombs(cell Cell) int {
	celldata := cell.GetData()

	fromRow := celldata.Row - 1
	toRow := celldata.Row + 1

	if fromRow < 0 {
		fromRow = 0
	}

	if toRow >= this.data.RowCount {
		toRow = this.data.RowCount - 1
	}

	fromCol := celldata.Col - 1
	toCol := celldata.Col + 1

	if fromCol < 0 {
		fromCol = 0
	}

	if toCol >= this.data.ColCount {
		toCol = this.data.ColCount - 1
	}

	count:=0
	for row := fromRow; row <= toRow; row++ {
		for col := fromCol; col <= toCol; col++ {
			currCell := this.getCell(row, col)
			if currCell.GetType() == shared.CellType_Bomb {
				count++
			}
		}
	}

	return count
}


