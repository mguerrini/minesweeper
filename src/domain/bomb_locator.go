package domain

import (
	"math/rand"
	"time"
)

type BombLocator interface {
	SetBombs(game *Game, countBombs int) error
}

type RandomBombLocator struct {

}

func NewRandomBombLocator() BombLocator {
	return &RandomBombLocator{}
}

func CreateRandomBombLocator (configurationName string) (interface{}, error) {
	return &RandomBombLocator{}, nil
}

func (this *RandomBombLocator) SetBombs(game *Game, countBombs int) error {
	rows, cols := game.GetRowCount(), game.GetColCount()

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	minesLeft := countBombs

	for minesLeft > 0 {
		bColNumber := rand.Intn(cols)
		bRowNumber := rand.Intn(rows)

		ok, err := game.SetBomb(bRowNumber, bColNumber)
		if err != nil {
			return err
		}

		if ok {
			minesLeft--
		}
	}

	return nil
}


type FixedBombLocator struct {

}

func NewFixedBombLocator() BombLocator {
	return &FixedBombLocator{}
}

func CreateFixedBombLocator (configurationName string) (interface{}, error) {
	return &FixedBombLocator{}, nil
}

func (this *FixedBombLocator) SetBombs(game *Game, countBombs int) error {
	return nil
}


