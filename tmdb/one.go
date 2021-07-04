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

func GetOne(c *gin.Context) {

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

	url := "https://api.themoviedb.org/3/movie/" + strconv.Itoa(id) + "?api_key=" + tmdb_key + "&language=en-US"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	resData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var resObj TDMovie
	json.Unmarshal(resData, &resObj)

	c.JSON(http.StatusOK, resObj)
}
