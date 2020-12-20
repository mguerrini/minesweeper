package domain

import "github.com/minesweeper/src/shared"

type Board struct {
	data shared.BoardData
	cells [][]Cell
}
