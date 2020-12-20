package domain

type MinesweeperGameFactory interface {
	CreateGame(rowCount, colCount int, bombCount int) (*Game, error)
}

type minesweeperGameFactory struct {
	bombLocator BombLocator
}
