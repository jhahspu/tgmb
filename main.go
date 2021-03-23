package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/data"
)

func main() {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(cors.Default())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "et voila",
			"mvs":     data.Rnd(),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	err := server.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
