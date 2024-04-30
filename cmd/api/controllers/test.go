package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RestHandler) Test(ctx *gin.Context) {
	msg := "test"
	data := "识别成功"
	ctx.JSON(http.StatusOK, &Response{
		Code:    SuccessCode,
		Message: msg,
		Data:    data,
	})
}
