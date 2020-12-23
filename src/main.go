package main

import (
	"github.com/minesweeper/src/common/logger"
	"github.com/minesweeper/src/server"
	"os"
)

func main(){
	//prepare server
	server.StartUp("local_configuration.yml")

	//start server
	engine := server.New()
	logger.Info("Listening and serving HTTP on port 8080")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	engine.Run(port)
}
