package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/db"
	"github.com/jhahspu/tgmb/tmdb"
)

func main() {

	if err := db.InitDBConnections(); err != nil {
		log.Fatalf("error connecting to db, %v \n", err)
	}

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(cors.Default())

	server.GET("/random", db.GetRandom)
	server.GET("/random/:genre", db.GetRandomByGenre)
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
