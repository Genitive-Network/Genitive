package app

import (
	"Genitive/cmd/api/routers"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	routers.RegisterRoutes(r)
	// live probe
	r.Run(":8006")
}
