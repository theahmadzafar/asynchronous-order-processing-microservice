package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func new(status int, data interface{}) *Response {
	response := &Response{
		Status: status,
		Data:   data,
	}

	if response.Data == nil {
		response.Data = http.StatusText(status)
	}

	if v, ok := data.(error); ok {
		response.Data = v.Error()
	}

	return response
}

func OK(ctx *gin.Context, data interface{}) {
	r := new(http.StatusOK, data)
	ctx.JSON(r.Status, data)
}

func Error(ctx *gin.Context, data interface{}) {
	r := new(http.StatusBadRequest, data)
	ctx.JSON(r.Status, data)
}
