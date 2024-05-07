package controllers

import (
	"Genitive/services"
	"github.com/gin-gonic/gin"
)

type GenitiveMintReq struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

func (r *RestHandler) Mint(ctx *gin.Context) {
	// business logic
	req := GenitiveMintReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	services.Mint(req.Address, req.Amount)

	var data = "成功"
	SuccessResponse(ctx, data, "Success")

}

type GenitiveBurnReq struct {
}

func (r *RestHandler) Burn(ctx *gin.Context) {
	req := GenitiveBurnReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	// business logic
	services.Burn()
	var data = ""
	SuccessResponse(ctx, data, "Success")
}
