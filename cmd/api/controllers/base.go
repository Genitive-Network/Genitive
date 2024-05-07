package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const APIV1 = "api/v1"

const SuccessCode = 200

type Response struct {
	Code      int         `json:"result"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"requestId"`
}

func SuccessResponse(ctx *gin.Context, data interface{}, Message ...string) {
	msg := "成功"
	if len(Message) > 0 {
		msg = Message[0]
	}
	ctx.JSON(http.StatusOK, &Response{
		Code:    SuccessCode,
		Message: msg,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, err error) {
	resp := &Response{
		Code:    1,
		Message: err.Error(),
		Data:    "",
	}
	c.JSON(http.StatusOK, resp)
}

func NotFoundResponse(ctx *gin.Context) {
	resp := &Response{
		Code:    1,
		Message: "您访问的地址不存在，请确认url是否正确",
	}
	ctx.JSON(http.StatusOK, resp)
}
