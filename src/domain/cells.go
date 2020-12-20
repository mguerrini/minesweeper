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

type cell struct {
	data shared.CellData
}

type bombCell struct {
	cell
}

type numberCell struct {
	cell
}




