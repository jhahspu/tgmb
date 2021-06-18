package models

type Movie struct {
	TmdbID      int    `json:"tmdb_id"`
	Title       string `json:"title"`
	Tagline     string `json:"tagline"`
	ReleaseDate string `json:"release_date"`
	Poster      string `json:"poster_path"`
	Backdrop    string `json:"backdrop_path"`
	Trailers    string `json:"trailers"`
}
