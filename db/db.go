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
	Runtime  string `json:"runtime"`
	Genres   string `json:"genres"`
	Overview string `json:"overview"`
	Poster   string `json:"poster_path"`
	Backdrop string `json:"backdrop_path"`
	Trailers string `json:"trailers"`
	Slug     string `json:"slug"`
}

const getRandom = `
SELECT tmdb, title, tagline, release, runtime, genres, overview, poster, backdrop, trailers, slug
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
			&i.Slug,
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
			&i.Slug,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	c.HTML(http.StatusOK, "random.tmpl", gin.H{
		"Title": "Random Titles",
		"rm":    items,
	})
}

const getRandomByGenre = `
SELECT tmdb, title, tagline, release, runtime, genres, overview, poster, backdrop, trailers, slug
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
			&i.Slug,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	c.JSON(http.StatusOK, items)
}

const getMovieBySlug = `
SELECT tmdb, title, tagline, release, runtime, genres, overview, poster, backdrop, trailers, slug
FROM mvs
WHERE slug=:slug
LIMIT 1
`

func GetMovieBySlug(c *gin.Context) {
	items := make([]ListMovies, 0)

	slug := c.Param("slug")

	rows, err := DBClient.Query(getMovieBySlug, slug)
	if err != nil {
		c.HTML(http.StatusBadRequest, "one.tmpl", gin.H{
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
			&i.Slug,
		); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	c.HTML(http.StatusOK, "one.tmpl", gin.H{
		"Title": "Random Titles",
		"rm":    items,
	})
}
