package factory



import "sync"

var (
	genericFactorySingleton GenericFactory
	genericFactoryOnce     sync.Once
)

func GenericFactorySingleton() GenericFactory {
	genericFactoryOnce.Do(func() {
		if genericFactorySingleton == nil {
			genericFactorySingleton = NewGenericFactory()
		}
	})

	return genericFactorySingleton;
}

func SetSingleton(inst GenericFactory) {
	if inst != nil {
		genericFactorySingleton = inst;
	}
}

