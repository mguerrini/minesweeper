package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/common/logger"
	"net/http"
)


func HandleResponse(c *gin.Context) error {
	c.Next()
	if len(c.Errors) > 0 {
		return nil
	}

	logRequestAndResponse(c)

	if response, exists := c.Get("response"); exists {
		statusCode := http.StatusOK
		if code, exists := c.Get("status_code"); exists {
			statusCode = code.(int)
		}

		c.JSON(statusCode, response)
	}

	return nil
}

func logRequestAndResponse(c *gin.Context) {
	//busco el parametro X-Traced
	var reqStr string
	var resStr string

	//REQUEST
	input, _ := c.Get("request")

	if input == nil {
		return
	}

	bsReq, errReq := json.Marshal(input);

	if errReq == nil {
		buf := bytes.NewBuffer(bsReq)
		reqStr = buf.String()
	}

	//RESPONSE
	output, _ := c.Get("response")

	if output != nil {
		bsRes, errRes := json.Marshal(output);

		if errRes == nil {
			buf := bytes.NewBuffer(bsRes)
			resStr = buf.String()
		}
	}

	//LOG
	logger.Info(fmt.Sprintf("{\"request\": %s, \"response\": %s }", reqStr, resStr))
}
