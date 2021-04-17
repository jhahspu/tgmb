package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jhahspu/tgmb/models/response"
)

type Trailers struct {
	ID      int       `json:"id"`
	Results []Trailer `json:"results"`
}

type Trailer struct {
	ID         string `json:"id"`
	ISO_639_1  string `json:"iso_639_1"`
	ISO_3166_1 string `json:"iso_3166_1"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	Site       string `json:"site"`
	Size       int    `json:"size"`
	Type       string `json:"type"`
}

func GetTrailers(c *gin.Context) {

	tmdb_key := os.Getenv("TMDB_KEY")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	url := "https://api.themoviedb.org/3/movie/" + strconv.Itoa(id) + "/videos?api_key=" + tmdb_key + "&language=en-US"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	resData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var resObj Trailers
	json.Unmarshal(resData, &resObj)

	c.JSON(http.StatusOK, resObj)
}
