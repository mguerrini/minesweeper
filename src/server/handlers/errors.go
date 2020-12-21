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

		//append all errors messages
		for _, ginError = range c.Errors {
			genericError, ok = ginError.Err.(apierrors.GenericError)
			errorsMsgs = append(errorsMsgs, ginError.Err.Error())
		}

		//find the generic error
		for _, ginError = range c.Errors {
			genericError, ok = ginError.Err.(apierrors.GenericError)
			if ok {
				break
			}
		}

		if genericError == nil {
			genericError = apierrors.NewApiError(ginError.Err, strings.Join(errorsMsgs, ","), http.StatusInternalServerError)
		}

		c.JSON(genericError.GetStatus(), genericError)

		logRequestWithError(c, genericError, errorsMsgs)

		if !c.IsAborted() {
			c.AbortWithStatus(genericError.GetStatus())
		}
	}
}

func logRequestWithError(c *gin.Context, mainErr apierrors.GenericError, msgList []string) {

	allMsgs := strings.Join(msgList, ",")

	input, _ := c.Get("request")

	if input != nil {
		bs, errSer := json.Marshal(input);

		if errSer == nil {
			buf := bytes.NewBuffer(bs)
			body := buf.String()
			logger.Error(fmt.Sprintf("Request Body: %s", body))
		}
	}

	logger.Error(fmt.Sprintf("%s - More errors: %s", mainErr.AsString(), allMsgs))
}
