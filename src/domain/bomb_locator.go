package domain

type BombLocator interface {
	SetBombs(game *Game, countBombs int)
}

type RandomBombLocator struct {

}

func NewRandomBombLocator() BombLocator {
	return &RandomBombLocator{}
}

func CreateRandomBombLocator (configurationName string) (interface{}, error) {
	return &RandomBombLocator{}, nil
}

func (this *RandomBombLocator) SetBombs(game *Game, countBombs int) {

}


type FixedBombLocator struct {

}

func NewFixedBombLocator() BombLocator {
	return &FixedBombLocator{}
}

func CreateFixedBombLocator (configurationName string) (interface{}, error) {
	return &FixedBombLocator{}, nil
}

func (this *FixedBombLocator) SetBombs(game *Game, countBombs int) {

}


