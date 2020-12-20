package factory


type FactoryRegistry interface {
	GetFactory(factoryName string) func(configurationName string) (interface{},  error)
	RegisterFactory(factoryName string, f func(configurationName string) (interface{},  error))
}

type factoryRegistry struct {
	factories map[string] func(configurationName string) (interface{}, error)
}


func CreateFactoryRegistry() FactoryRegistry  {
	return &factoryRegistry{
		factories: make(map[string]func(configurationName string) (interface{}, error)),
	}
}

func (this *factoryRegistry) RegisterFactory(factoryName string, f func(configurationName string) (interface{},  error)) {
	this.factories[factoryName] = f
}

func (this *factoryRegistry) GetFactory(factoryName string) func(configurationName string) (interface{},  error) {
	if fact, ok := this.factories[factoryName]; ok {
		return fact;
	}

	return nil
}
