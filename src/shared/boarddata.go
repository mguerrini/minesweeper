package shared

type BoardData struct {
	RowCount   int          `json:"row_count"`
	ColCount   int          `json:"col_count"`
	BombsCount int          `json:"bomb_count"`
	Cells      [][]CellData `json:"cells"`
}


func (this *BoardData) GetCell(row int, col int) CellData {
	return this.Cells[row][col]
}

func (this *BoardData) SetCell(row int, col int, cell CellData)  {
	this.Cells[row][col] = cell
}

//Clean the data of the cells that have not been exposed
func (this *BoardData) HideNotRevealed() {
	for row:=0;row <this.RowCount; row++ {
		for col:=0;col <this.ColCount; col++ {
			c := this.GetCell(row, col)
			c.Hide()
			this.SetCell(row, col, c)
		}
	}
}

func (this *BoardData) RevealAll() {
	for row:=0;row <this.RowCount; row++ {
		for col:=0;col <this.ColCount; col++ {
			c := this.GetCell(row, col)
			c.Reveal()
			this.SetCell(row, col, c)
		}
	}
}

