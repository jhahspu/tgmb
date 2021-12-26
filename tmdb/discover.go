package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/models/response"

	"github.com/joho/godotenv"
)

type Discover struct {
	Page    int       `json:"page"`
	Results []TDMovie `json:"results"`
}

type TDMovie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        float32 `json:"vote_count"`
}

func GetDiscover(c *gin.Context) {

	tmdb_key := os.Getenv("TMDB_KEY")

	url := "https://api.themoviedb.org/3/discover/movie?api_key=" + tmdb_key + "&language=en-US&sort_by=popularity.desc&include_adult=false&include_video=true&page=1&year=2021"

	httpResponse, err := http.Get(url)
	if err != nil {
		res := response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	resData, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(err)
	}

	var resObj Discover
	json.Unmarshal(resData, &resObj)

	c.JSON(http.StatusOK, resObj)
}

func DiscoverPage(c *gin.Context) {

	godotenv.Load(".env")
	tmdb_key := os.Getenv("TMDB_KEY")

	url := "https://api.themoviedb.org/3/discover/movie?api_key=" + tmdb_key + "&language=en-US&sort_by=popularity.desc&include_adult=false&include_video=true&page=1&year=2021"

	httpResponse, err := http.Get(url)
	if err != nil {
		res := response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	resData, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(err)
	}

	var resObj Discover
	json.Unmarshal(resData, &resObj)

	c.HTML(http.StatusOK, "discover.tmpl", gin.H{
		"Title": "Discover",
		"rm":    resObj.Results,
	})

}
