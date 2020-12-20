package configs

func CreateConfigurationManager (path string, filename string) ConfigurationManager {
	output := &configurationManager {}

	output.Load(path, filename)

	return output
}


