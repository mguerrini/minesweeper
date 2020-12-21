package server

import (
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/common/configs"
	"github.com/minesweeper/src/controllers"
	"github.com/minesweeper/src/services"
)

var (
	Router *gin.Engine
	Controller *controllers.MinesweeperController
)


func StartUp() {
	configs.Initialize("local_configuration.yml")

	RegisterFactories()

	Controller = controllers.NewMinesweeperController()

	//search service configuration name
	if !configs.Singleton().Exist("root.startup.minesweeper") ||
		configs.Singleton().IsNil("root.startup.minesweeper") {
		panic("MinesweeperService configuration name not exist or is not defined in path: root.startup.minesweeper")
	}

	serviceFactoryConfigName, _ :=  configs.Singleton().GetString("root.startup.minesweeper")

	service, err := services.NewMinesweeperService(serviceFactoryConfigName)

	if err != nil {
		panic(err)
	}

	//set singleton instance
	services.SetSingleton(service)
}
