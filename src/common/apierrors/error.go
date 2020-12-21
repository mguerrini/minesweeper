package apierrors

import (
	"errors"
	"fmt"
	"net/http"
)

type GenericError interface {
	GetStatus() int
	Error() string
	Cause() error
	ErrorClass() string
	AsString() string
}

type ApiError struct {
	code    int    `json:"code"`
	status  string `json:"status"`
	error   string `json:"error"`
	message string `json:"message"`
}


func NewBadRequest(err error, message string) *ApiError {
	return NewApiError(err, message, http.StatusBadRequest)
}


func NewInternalServerError(err error, message string) *ApiError {
	return NewApiError(err, message, http.StatusInternalServerError)
}

func NewApiError(err error, message string, status int) *ApiError {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return &ApiError{
		message: message,
		error:   errMsg,
		status:  http.StatusText(status),
		code:    status,
	}
}


func (e *ApiError) GetMessage() string {
	return e.message
}

func (e *ApiError) GetStatusCode() int {
	return e.code
}

func (e *ApiError) GetError() string {
	return e.error
}

func (e *ApiError) AsString() string {
	if len(e.error) == 0 {
		if len(e.GetMessage()) == 0 {
			return fmt.Sprintf("{[code: %s - %s]}", e.code, e.status, e.GetMessage())
		} else {
			return fmt.Sprintf("{[code: %s - %s] message: %s}", e.code, e.status, e.GetMessage())
		}
	} else {
		if len(e.GetMessage()) == 0 {
			return fmt.Sprintf("{[code: %s - %s] error: %s}", e.code, e.status, e.GetMessage(), e.GetError())
		} else {
			return fmt.Sprintf("{[code: %s - %s] message: %s, error: %s}", e.code, e.status, e.GetMessage(), e.GetError())
		}
	}
}

/*************************************/
/*      GenericError interface       */
/*************************************/
func (e *ApiError) GetStatus() int {
	return e.code
}

func (e *ApiError) Error() string {
	return e.error
}

func (e *ApiError) Cause() error {
	if len(e.error) == 0 {
		return errors.New(e.message)
	} else {
		return errors.New(e.error)
	}
}

func (e *ApiError) ErrorClass() string {
	return e.status
}

/*************************************/
