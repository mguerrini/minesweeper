package server

import (
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/controllers"
	"github.com/minesweeper/src/server/handlers"
	"net/http"
)

var (
	Router *gin.Engine
	Controller *controllers.MinesweeperController
)

// New creates a new router
func New() *gin.Engine {

	configureRouter()

	Controller = CreateController()

	return mapUrls(Router)
}

func configureRouter() {
	gin.SetMode(gin.ReleaseMode)
	// Default meli router - includes newrelic, datadog, attributes filter, jsonp and pprof middlewares
	Router = gin.New()

	// Panic recovery
	Router.Use(handlers.RecoveryWithWriter())
	Router.Use(handlers.HandleError)
	Router.Use(AdaptHandler(handlers.HandleResponse))
}

func mapUrls(router *gin.Engine) *gin.Engine {

	router.GET("/ping",
		AdaptHandler(Pong))

	router.GET("minesweeper/users/:user_id/games",
		AdaptHandler(Controller.GetGameListByUserId))

	router.GET("minesweeper/users/:user_id/games/:game_id",
		AdaptHandler(Controller.GetGame))

	router.POST("minesweeper/users/:user_id/games",
		AdaptHandler(Controller.CreateNewGame))

	router.POST("minesweeper/users/:user_id/games/:game_id/reveal",
		AdaptHandler(Controller.RevealCell))

	router.POST("minesweeper/users/:user_id/games/:game_id/mark",
		AdaptHandler(Controller.MarkCell))

	router.DELETE("minesweeper/users/:user_id/games/:game_id",
		AdaptHandler(Controller.DeleteGame))

	router.DELETE("minesweeper/users/:user_id/games",
		AdaptHandler(Controller.DeleteGamesByUser))

	return router
}

//GET ping
func Pong(c *gin.Context) error {
	c.Set("skip", true)
	c.JSON(http.StatusOK, "Pong from minesweeper")
	return nil
}