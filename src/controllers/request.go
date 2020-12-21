package controllers

type NewGameRequest struct {
	Rows    int    `json:"rows"`
	Columns int    `json:"columns"`
	Mines   int    `json:"mines"`
}

type BaseGameActionRequest struct {
	Row    int    `json:"row"`
	Col    int    `json:"col"`
}

type MarkCellRequest struct {
	BaseGameActionRequest
	None bool `json:"none"`
	Flag bool `json:"flag"`
	Question bool `json:"question"`
}

type RevealCellRequest struct {
	BaseGameActionRequest
}


