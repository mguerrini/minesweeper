package factory

import "sync"

var (
	factoryRegistrySingleton FactoryRegistry
	factoryRegistryOnce     sync.Once
)

func FactoryRegistrySingleton() FactoryRegistry {
	factoryRegistryOnce.Do(func() {
		if factoryRegistrySingleton == nil {
			factoryRegistrySingleton = CreateFactoryRegistry()
		}
	})

	return factoryRegistrySingleton;
}

func SetFactoryRegistrySingleton(inst FactoryRegistry) {
	if inst != nil {
		factoryRegistrySingleton = inst;
	}
}
