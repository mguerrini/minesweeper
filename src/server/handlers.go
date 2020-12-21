package server

import (
	"github.com/gin-gonic/gin"
)


func AdaptHandler(handlerFn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := handlerFn(c); err != nil {
			_ = c.Error(err)
			c.Abort()
		}
	}
}

