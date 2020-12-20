package domain

import (
	"github.com/minesweeper/src/shared"
)

type Game struct {
	data shared.GameData
	board Board
}


func (this *Game) GetId() string {
	return this.data.Id
}

func (this *Game) SetId(id string)  {
	this.data.Id = id
}


