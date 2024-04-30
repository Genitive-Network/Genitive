package routers

import (
	"Genitive/cmd/api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	RestHandler := controllers.NewHandler()

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	router.POST("/api/mint", RestHandler.Mint)
	router.POST("/api/burn", RestHandler.Burn)
}
