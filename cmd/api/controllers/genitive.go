package controllers

import (
	"github.com/gin-gonic/gin"
)

type GenitiveReq struct {
}

func (r *RestHandler) Mint(ctx *gin.Context) {
	// business logic
	req := GenitiveReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	var data = ""
	SuccessResponse(ctx, data, "Success")

}

func (r *RestHandler) Burn(ctx *gin.Context) {
	// business logic
}
