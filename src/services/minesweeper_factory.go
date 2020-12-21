package services

import (
	"github.com/minesweeper/src/common/configs"
	"github.com/minesweeper/src/common/factory"
	"github.com/minesweeper/src/dal/gamedal"
	"github.com/minesweeper/src/domain"
)

type MinesweeperServiceConfiguration struct {
	MinesLocatorConfigurationName string `json:"mineslocator"`
	GameDalConfigurationName      string `json:"gamedal"`
}


//Factory Method
func CreateMinesweeperService (configurationName string) (interface{}, error) {
	var minesLocator domain.MinesLocator
	var gameDal gamedal.GameDal

	if len(configurationName) == 0 || !configs.Singleton().Exist(configurationName) {
		return createDefaultMinesweeperService()
	} else {
		//search configuration
		conf := &MinesweeperServiceConfiguration{}
		err := configs.Singleton().GetObject(configurationName, conf)
		if err != nil {
			return nil, err
		}

		if len(conf.MinesLocatorConfigurationName) == 0 {
			minesLocator = domain.NewRandomMinesLocator()
		} else {
			minesLocObj, err :=	factory.GenericFactorySingleton().Create(conf.MinesLocatorConfigurationName)
			if err != nil {
				return nil, err
			}
			minesLocator = minesLocObj.(domain.MinesLocator)
		}

		if len(conf.GameDalConfigurationName) == 0 {
			gameDal, err = gamedal.NewDbGameDal("")
			if (err != nil) {
				return nil, err
			}
		} else {
			gameDalObj, err :=	factory.GenericFactorySingleton().Create(conf.GameDalConfigurationName)
			if err != nil {
				return nil, err
			}
			gameDal = gameDalObj.(gamedal.GameDal)
		}
	}

	//set gamefactory and dal
	output := &minesweeperService {}
	output.gameFactory = domain.NewMinesweeperGameFactory(minesLocator)
	output.gameDal = gameDal

	return output, nil
}

func NewMinesweeperService(factoryConfigurationName string) (MinesweeperService, error) {
	if len(factoryConfigurationName) == 0 {
		output, err := CreateMinesweeperService("")
		return output.(MinesweeperService), err
	}

	instance, err := factory.GenericFactorySingleton().Create(factoryConfigurationName)

	if err != nil {
		return nil, err
	}

	if instance != nil {
		return instance.(MinesweeperService), nil
	}

	//return with default configuration
	return createDefaultMinesweeperService()
}

func createDefaultMinesweeperService() (MinesweeperService, error){
	//create with defualt values

	gameDal, err := gamedal.NewDbGameDal("")
	if (err != nil) {
		return nil, err
	}
	minesLocator := domain.NewRandomMinesLocator()

	//set gamefactory and dal
	output := &minesweeperService {}
	output.gameFactory = domain.NewMinesweeperGameFactory(minesLocator)
	output.gameDal = gameDal

	return output, nil
}

