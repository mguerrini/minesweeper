package domain

type BombLocator interface {
	SetBombs(game *Game, countBombs int)
}

type RandomBombLocator struct {

}

type FixedBombLocator struct {

}
