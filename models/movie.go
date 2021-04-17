package models

type Movie struct {
	TmdbID      int    `json:"tmdb_id"`
	Title       string `json:"title"`
	Tagline     string `json:"tagline"`
	ReleaseDate string `json:"release_date"`
	Backdrop    string `json:"backdrop"`
	Trailers    string `json:"trailers"`
}
