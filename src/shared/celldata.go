package shared

type CellData struct {
	Type      CellType `json:"type"`
	Row       int      `json:"row"`
	Col       int      `json:"col"`
	IsReveled bool     `json:"is_exposed"`
	IsMarked  bool     `json:"is_marked"`
	Number    int      `json:"number"`
}

