package factory

import (
	"errors"
	"github.com/minesweeper/src/common/configs"
)

type GenericFactory interface {
	Create(factoryConfigurationName string) (interface{},  error)
}

type genericfactory struct {
}

func NewGenericFactory() GenericFactory {
	return &genericfactory{}
}

func (this *genericfactory) Create(factoryConfigurationName string) (interface{}, error) {
	if len(factoryConfigurationName) == 0 {
		return nil, errors.New("The name of the configuration can not be empty")
	}

	factoryNameRef := factoryConfigurationName + ".factory"

	if !configs.Singleton().Exist(factoryNameRef) ||
		configs.Singleton().IsNil(factoryNameRef) {
		return nil, nil
	}


	//get the factory name to use
	factoryName, _ := configs.Singleton().GetString(factoryNameRef)

	fact := FactoryRegistrySingleton().GetFactory(factoryName)

	if fact == nil {
		return nil, nil
	}

	//create de instance
	confSectionName := factoryConfigurationName + ".configuration"
	output, err := fact(confSectionName)

	if err != nil {
		return nil, err
	}

	return output, nil
}

