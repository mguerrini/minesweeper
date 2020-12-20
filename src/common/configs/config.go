package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/minesweeper/src/common/helpers"
	"github.com/olebedev/config"
	"os"
	folder "path"
)

type ConfigurationManager interface {
	Exist(path string) bool
	IsNil(path string)  bool
	Clean()
	GetObject(path string, emptyObj interface{}) error
	GetString(path string) (string, error)
}


type configurationManager struct {
	cfg *config.Config
}

type configEnvelope struct {
	Root interface{} `json:"Root"`
}


func (this *configurationManager) GetBasePath() string {
	return os.Getenv("GOPATH")
}

func (this *configurationManager) GetConfigPath() string {
	confDir := os.Getenv("CONF_DIR")
	if confDir == "" {
		base := os.Getenv("GOPATH")
		confDir = folder.Join(base, "src/github.com/minesweeper/configs")
		fmt.Fprintf(os.Stdout, "CONF_DIR env variable was not set, use default: "+ confDir +"\n")
	}

	return confDir
}


//Load a file and join it with thw existent configuration
func (this *configurationManager) Load(path string, file string) {
	cfg := this.doParseFile(path, file)
	this.cfg = cfg
}

//Load a new configuration file.
//If exist a previous configuration loaded, is removed and replaced wuth the new one.
func (this *configurationManager) Join(path string, file string) {
	toJoinConf := this.doParseFile(path, file)
	nCfg, err := this.cfg.Extend(toJoinConf)

	if err != nil {
		panic(err.Error())
	}

	this.cfg = nCfg
	fmt.Fprintf(os.Stdout, "Configuration file (%s) loaded", file)
}

func (this *configurationManager) doParseFile(path string, file string) *config.Config {

	file = this.validateFile(path, file)

	fmt.Fprintf(os.Stdout, "Reading configuration from: %s\n", file)

	var err error
	cfg, err := config.ParseYamlFile(file)
	if err != nil {
		panic(err.Error())
	}

	return cfg
}

func (this *configurationManager) validateFile(path string, file string) string {

	if len(path) == 0 {
		var err error

		//try on configuration folder
		path = this.GetConfigPath()
		fileAux1 := folder.Join(path, file)

		if this.fileExists(fileAux1) {
			return fileAux1
		}

		//try in base go path
		//path = this.GetBasePath()
		fileAux2 := folder.Join(path, file)
		if this.fileExists(fileAux2) {
			return fileAux2
		}

		//try on working folder
		wd, wderr := os.Getwd()

		if wderr != nil {
			panic(err)
		}

		fileAux3 := folder.Join(path, wd)
		if this.fileExists(fileAux3) {
			return fileAux3
		}

		panic(errors.New("The configuracion file " + file + " not exist. Look in: "  + fileAux1 + " - " + fileAux2 + " - " + fileAux3))
	}

	//uso el path
	fileAux := folder.Join(path, file)
	if _, err4 := os.Stat(file); os.IsNotExist(err4) {
		panic(err4)
	}

	return fileAux
}


func (this *configurationManager) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (this *configurationManager) Clean() {
	this.cfg = nil
}

func (this *configurationManager) Exist(path string)  (bool) {
	if this.cfg == nil {
		return false
	}
	_, err := this.cfg.Get(path)

	return err == nil
}



func (this *configurationManager) IsNil(path string)  bool {
	if this.cfg == nil {
		return true
	}

	val, err := this.cfg.Get(path)

	if err != nil {
		return true
	}

	return val.Root == nil
}


func (this *configurationManager) GetString(path string)  (string, error) {
	if this.cfg == nil {
		return "", errors.New("The configuration is not loaded")
	}

	return this.cfg.String(path)
}


func (this *configurationManager) GetObject(path string, configType interface{})  error {
	if this.cfg == nil {
		return errors.New("The configuration is not loaded")
	}


	newConfig, err := this.cfg.Get(path)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error getting path to configuration. Path: %s - ConfigType: %s - Error: %s",
			path,
			helpers.GetTypeName(configType),
			err.Error(),
		)
		return err
	}

	jsonObj, err := config.RenderJson(newConfig)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error getting json data for configuration. Path: %s - ConfigType: %s - Error: %s",
			path,
			helpers.GetTypeName(configType),
			err.Error(),
		)
		return err
	}

	//json to configs
	objBytes := []byte(jsonObj)

	env := configEnvelope{Root: configType}
	err = json.Unmarshal(objBytes, &env)

	if err != nil {
		return err
	}

	configType = env.Root

	return nil
}

