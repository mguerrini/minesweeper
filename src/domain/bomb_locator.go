package domain

import (
	"math/rand"
	"time"
)

type BombLocator interface {
	SetBombs(game *Game, countBombs int) error
}

/***********************************/
/*      Random Bomb Locator        */
/***********************************/

type RandomBombLocator struct {

}

func NewRandomBombLocator() *RandomBombLocator {
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

/***********************************/
/*       Fixed Bomb Locator        */
/***********************************/

type FixedBombLocator struct {
 	BombsCoordinates []boadCoordinate
}

type fixedBombLocatorConfiguration struct {
	bombs []string `json:"bombs"`
}

type boadCoordinate struct {
	Row int
	Col int
}

func NewFixedBombLocator() *FixedBombLocator {
	return &FixedBombLocator{
		BombsCoordinates: make([]boadCoordinate, 0),
	}
}

func CreateFixedBombLocator (configurationName string) (interface{}, error) {
	//get configuration

	return NewFixedBombLocator(), nil
}

func (this *FixedBombLocator) AddBomb(row, col int)  {
	this.BombsCoordinates = append(this.BombsCoordinates, boadCoordinate{
		Row: row,
		Col: col,
	})
}

func (this *FixedBombLocator) SetBombs(game *Game, countBombs int) error {
	for _, c := range this.BombsCoordinates {
		_, err := game.SetBomb(c.Row, c.Col)

		if err != nil {
			return err
		}
	}

	return nil
}


