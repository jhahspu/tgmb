package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/data"
)

func main() {
	server := gin.Default()

	server.GET("/rnd", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "et voila",
			"mvs":     data.Rnd(),
		})
	})

	server.Run(":9000")
}
