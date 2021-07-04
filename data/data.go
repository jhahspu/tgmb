package data

import (
	"encoding/csv"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
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
	sort.Ints(rs)
	mvs := make([]models.Movie, 0, 20)
	wg := sync.WaitGroup{}
	for _, pos := range rs {
		wg.Add(1)
		go func(pos int) {
			mvs = append(mvs, getMovie(pos, records))
			wg.Done()
		}(pos)
	}
	wg.Wait()
	c.JSON(http.StatusOK, mvs)
}

/**
*		Find Row by ID
*			& return Movie
**/
func getMovie(id int, records [][]string) (movie models.Movie) {
	mv := models.Movie{}
	for i, record := range records {
		if i == id {
			mv.TmdbID, _ = strconv.Atoi(record[1])
			mv.Title = record[2]
			mv.Tagline = record[3]
			mv.ReleaseDate = record[4]
			mv.Poster = record[5]
			mv.Backdrop = record[6]
			mv.Trailers = record[7]
		}
	}
	return mv
}

/**
*		Create Random Slice of N intergers
*		Return first 120
*
**/
func randSlice(n int) []int {
	rand.Seed(time.Now().Unix())
	x := rand.Perm(n)
	y := x[:20]
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
