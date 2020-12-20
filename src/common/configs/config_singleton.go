package configs

import "sync"

var (
	singleton ConfigurationManager
	once     sync.Once
)

func Singleton() ConfigurationManager {
	once.Do(func() {
		if singleton == nil {
			//TODO Obtener variables de entorno para determinar cual cargar
			singleton = CreateConfigurationManager("", "local_configuration.yml")
		}
	})

	return singleton;
}

func SetSingleton(inst ConfigurationManager) {
	if inst != nil {
		singleton = inst;
	}
}

func Initialize(file string) {
	inst := CreateConfigurationManager("", file)
	SetSingleton(inst)
}


