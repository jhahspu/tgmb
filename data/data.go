package data

import (
	"encoding/csv"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Movie struct {
	TmdbID      int    `json:"tmdb_id"`
	Title       string `json:"title"`
	Tagline     string `json:"tagline"`
	ReleaseDate string `json:"release_date"`
	Backdrop    string `json:"backdrop"`
	Trailers    string `json:"trailers"`
}

func Rnd() []Movie {
	url := "https://raw.githubusercontent.com/jhahspu/tgm_vuejs/main/resources/movies_1616.csv"
	fm, err := http.Get(url)
	if err != nil {
		log.Fatalf("unable to read from url, %v", err)
	}
	defer fm.Body.Close()

	cv := csv.NewReader(fm.Body)
	records, err := cv.ReadAll()
	if err != nil {
		log.Fatalf("unable to parse file as CSV, %v", err)
	}

	rs := randSlice(1616)

	mvs := make([]Movie, 0, 200)
	for _, pos := range rs {
		for i, record := range records[1:] {
			if i == pos {
				mv := Movie{}
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
	return mvs
}

func randSlice(n int) []int {
	rand.Seed(time.Now().Unix())
	x := rand.Perm(n)
	y := x[:120]
	return y
}
