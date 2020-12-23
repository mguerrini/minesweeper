package main

import (
	"github.com/minesweeper/src/common/logger"
	"github.com/minesweeper/src/server"
)

func main(){
	//prepare server
	server.StartUp("local_configuration.yml")

	//start server
	engine := server.New()
	logger.Info("Listening and serving HTTP on port 8080")

	engine.Run(":8080")
}
