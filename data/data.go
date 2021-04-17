package data

import (
	"encoding/csv"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/models"
)

/**
*		Rnd will return 120 movies
*
**/
func Rnd(c *gin.Context) {
	records := readCSV("movies.csv")
	rs := randSlice(1616)
	mvs := make([]models.Movie, 0, 200)
	for _, pos := range rs {
		for i, record := range records[1:] {
			if i == pos {
				mv := models.Movie{}
				mv.TmdbID, _ = strconv.Atoi(record[1])
				mv.Title = record[2]
				mv.Tagline = record[3]
				mv.ReleaseDate = record[4]
				mv.Backdrop = record[5]
				mv.Trailers = record[6]
				mvs = append(mvs, mv)
			}
		}
	}

	c.JSON(http.StatusOK, mvs)
}

/**
*		Create Random Slice of N intergers
*		Return first 120
*
**/
func randSlice(n int) []int {
	rand.Seed(time.Now().Unix())
	x := rand.Perm(n)
	y := x[:120]
	return y
}

/**
*		Read CSV from file
*		Return records
**/
func readCSV(filepath string) [][]string {
	f, err := os.Open(filepath)
	handleErrors(err, "Unable to read input file")
	defer f.Close()
	csvreader := csv.NewReader(f)
	records, err := csvreader.ReadAll()
	handleErrors(err, "Unable to parse file as CSV")
	return records
}

/**
*		Error Handler
*
**/
func handleErrors(err error, msg string) {
	if err != nil {
		log.Fatalf("[Error] %s %v", msg, err)
	}
}
