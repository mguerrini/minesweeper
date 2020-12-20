package services

import "sync"

var (
	singleton MinesweeperService
	once     sync.Once
)

func Singleton() MinesweeperService {
	once.Do(func() {
		if singleton == nil {
			s, err := NewMinesweeperService("minesweeper.services.minesweeper.default")

			if err != nil {
				panic(err)
			}

			singleton = s
		}
	})

	return singleton
}

func SetSingleton(inst MinesweeperService) {
	if inst != nil {
		singleton = inst;
	}
}


