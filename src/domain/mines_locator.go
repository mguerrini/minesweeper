package domain

import (
	"github.com/minesweeper/src/common/configs"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type MinesLocator interface {
	SetMines(game *Game, countMines int) error
}

/***********************************/
/*      Random Mines Locator        */
/***********************************/

type RandomMinesLocator struct {

}

func NewRandomMinesLocator() *RandomMinesLocator {
	return &RandomMinesLocator{}
}

func CreateRandomMinesLocator(configurationName string) (interface{}, error) {
	return &RandomMinesLocator{}, nil
}

func (this *RandomMinesLocator) SetMines(game *Game, minesCount int) error {
	rows, cols := game.GetRowCount(), game.GetColCount()

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	minesLeft := minesCount

	for minesLeft > 0 {
		bColNumber := rand.Intn(cols)
		bRowNumber := rand.Intn(rows)

		ok, err := game.SetMines(bRowNumber, bColNumber)
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
/*       Fixed Mines Locator        */
/***********************************/

type FixedMineLocator struct {
 	MinesCoordinates []boardCoordinate
}

type fixedMinesLocatorConfiguration struct {
	Mines []string `json:"mines"`
}

type boardCoordinate struct {
	Row int
	Col int
}

func NewFixedMineLocator() *FixedMineLocator {
	return &FixedMineLocator{
		MinesCoordinates: make([]boardCoordinate, 0),
	}
}

func CreateFixedMinesLocator(configurationName string) (interface{}, error) {
	//get configuration
	conf := fixedMinesLocatorConfiguration{}
	configs.Singleton().GetObject(configurationName, &conf)

	output := NewFixedMineLocator()
	if conf.Mines != nil {
		for _, valStr := range conf.Mines {
			separator := strings.Index(valStr, ",")
			rowStr := valStr[:separator]
			colStr := valStr[separator+1:]

			row, err := strconv.Atoi(rowStr)
			if err != nil {
				return nil, err
			}
			col, err := strconv.Atoi(colStr)
			if err != nil {
				return nil, err
			}

			output.AddMine(row, col)
		}
	}

	return output, nil
}

func (this *FixedMineLocator) AddMine(row, col int)  {
	this.MinesCoordinates = append(this.MinesCoordinates, boardCoordinate{
		Row: row,
		Col: col,
	})
}

func (this *FixedMineLocator) SetMines(game *Game, minesCount int) error {
	for _, c := range this.MinesCoordinates {
		_, err := game.SetMines(c.Row, c.Col)

		if err != nil {
			return err
		}
	}

	return nil
}


