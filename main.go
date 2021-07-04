package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/data"
	"github.com/jhahspu/tgmb/tmdb"
)

func main() {

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(cors.Default())

	server.GET("/random", data.Rnd)
	server.GET("/discover", tmdb.GetDiscover)
	server.GET("/trailers/:id", tmdb.GetTrailers)
	server.GET("/details/:id", tmdb.GetOne)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	err := server.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
