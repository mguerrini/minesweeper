package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/common/apierrors"
	"github.com/minesweeper/src/common/logger"
	"net/http"
	"strings"
)

func HandleError(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		var errorsMsgs = make([]string, 0, len(c.Errors))
		var ginError *gin.Error
		var genericError apierrors.GenericError
		var ok bool
		for _, ginError = range c.Errors {
			errorsMsgs = append(errorsMsgs, ginError.Err.Error())
			genericError, ok = ginError.Err.(apierrors.GenericError)
			if !ok {
				genericError = apierrors.NewApiError(ginError.Err, strings.Join(errorsMsgs, ","), http.StatusInternalServerError)
			}
		}

		if genericError != nil {
			logger.Error(genericError.AsString())

			c.JSON(genericError.GetStatus(), genericError)

			logRequestWithError(c, genericError)
		} else {
			panic("no error found when expected")
		}

		if !c.IsAborted() {
			c.AbortWithStatus(genericError.GetStatus())
		}
	}
}

func logRequestWithError(c *gin.Context, err error) {
	input, _ := c.Get("request")

	if input != nil {
		bs, errSer := json.Marshal(input);

		if errSer == nil {
			buf := bytes.NewBuffer(bs)
			body := buf.String()
			logger.Error(fmt.Sprintf("Request Body: %s\nError: %s", body, err.Error()))
		}
	}
}
