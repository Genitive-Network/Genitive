package controllers

import (
	"Genitive/services"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type GenitiveMintReq struct {
	Address string  `json:"address" binding:"required"`
	Amount  big.Int `json:"amount" binding:"required"`
}

func (r *RestHandler) Mint(ctx *gin.Context) {
	// business logic
	req := GenitiveMintReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	log.Println("req", req.Address, req.Amount)

	var data = StatusResp{Status: false}
	err := services.Mint(req.Address, req.Amount)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	data.Status = true
	SuccessResponse(ctx, data, "Success")

}

type GenitiveBurnReq struct {
	Address string          `json:"address" binding:"required"`
	Amount  decimal.Decimal `json:"amount" binding:"required"`
}

func (r *RestHandler) Burn(ctx *gin.Context) {
	req := GenitiveBurnReq{}
	var data = StatusResp{Status: false}
	if err := ctx.ShouldBind(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	// business logic
	err := services.Burn(req.Address, req.Amount)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
	data.Status = true

	SuccessResponse(ctx, data, "Success")
}

type StatusResp struct {
	Status bool `json:"status"`
}
