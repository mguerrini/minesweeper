package main

import (
	"github.com/minesweeper/src/common/configs"
)

func main(){
	//TODO Check if local or production environment
	configs.Initialize("local_configuration.yml")
}
