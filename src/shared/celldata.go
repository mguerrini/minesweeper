package shared

type CellData struct {
	Type       CellType     `json:"type"`
	Row        int          `json:"row"`
	Col        int          `json:"col"`
	IsRevealed bool         `json:"is_revealed"`
	Mark       CellMarkType `json:"mark"`
	Number     int          `json:"number"`
}


func (this *CellData) Hide() {
	if this.IsRevealed {
		return
	}

	this.Type = CellType_Unknown
	this.Number =-1
}

func (this *CellData) Reveal() {
	this.IsRevealed = true;
	this.Mark = CellMarkType_None
}
