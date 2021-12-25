package db

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DBClient *sql.DB

func InitDBConnections() error {
	db, err := sql.Open("sqlite3", "./db/tgm.db")
	if err != nil {
		return err
	}
	DBClient = db
	return nil
}

type ListMovies struct {
	TMDb     int    `json:"id"`
	Title    string `json:"title"`
	Tagline  string `json:"tagline"`
	Release  string `json:"release_date"`
	Runtime  int    `json:"runtime"`
	Genres   string `json:"genres"`
	Overview string `json:"overview"`
	Poster   string `json:"poster_path"`
	Backdrop string `json:"backdrop_path"`
	Trailers string `json:"trailers"`
}

const getRandom = `
SELECT tmdb, title, tagline, release, runtime, genres, overview, poster, backdrop, trailers
FROM mvs
ORDER BY random()
LIMIT 20
`

func GetRandom(c *gin.Context) {
	items := make([]ListMovies, 0)

	rows, err := DBClient.Query(getRandom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i ListMovies
		if err := rows.Scan(
			&i.TMDb,
			&i.Title,
			&i.Tagline,
			&i.Release,
			&i.Runtime,
			&i.Genres,
			&i.Overview,
			&i.Poster,
			&i.Backdrop,
			&i.Trailers,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	c.JSON(http.StatusOK, items)
}

func RandomPage(c *gin.Context) {
	items := make([]ListMovies, 0)

	rows, err := DBClient.Query(getRandom)
	if err != nil {
		c.HTML(http.StatusBadRequest, "random.tmpl", gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i ListMovies
		if err := rows.Scan(
			&i.TMDb,
			&i.Title,
			&i.Tagline,
			&i.Release,
			&i.Runtime,
			&i.Genres,
			&i.Overview,
			&i.Poster,
			&i.Backdrop,
			&i.Trailers,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	c.HTML(http.StatusOK, "random.tmpl", gin.H{
		"rm": items,
	})
}

const getRandomByGenre = `
SELECT tmdb, title, tagline, release, runtime, genres, overview, poster, backdrop, trailers
FROM mvs
WHERE instr(genres, ?)
ORDER BY random()
LIMIT 20
`

func GetRandomByGenre(c *gin.Context) {
	items := make([]ListMovies, 0)

	genre := c.Param("genre")

	rows, err := DBClient.Query(getRandomByGenre, genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i ListMovies
		if err := rows.Scan(
			&i.TMDb,
			&i.Title,
			&i.Tagline,
			&i.Release,
			&i.Runtime,
			&i.Genres,
			&i.Overview,
			&i.Poster,
			&i.Backdrop,
			&i.Trailers,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	c.JSON(http.StatusOK, items)
}
