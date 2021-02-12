package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minesweeper/src/common/logger"
	"io/ioutil"
	"net/http"
)


func HandleRequestAndResponse(c *gin.Context) error {
	buf := make([]byte, 1024)
	num, _:= c.Request.Body.Read(buf)
	c.Request.Body.Close()

	req := string(buf[0:num])
	c.Set("request", req)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf[0:num]))

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

	var reqMap map[string]interface{}
	reqStr, _ = input.(string) // json.Marshal(input);
	json.Unmarshal([]byte(reqStr), &reqMap);
	bsReq, _ := json.Marshal(reqMap)
	reqStr = string(bsReq)
/*
	bsReq, _ := input.(string) // json.Marshal(input);
	bsReq = strings.Replace(bsReq, "\n", "", -1)
	reqStr = bsReq
*/
/*
	if errReq == nil {
		buf := bytes.NewBuffer(bsReq)
		reqStr = buf.String()
	}
*/

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
	logger.Info(fmt.Sprintf("Payloads: {\"request\": %s, \"response\": %s }", reqStr, resStr))
}
