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

