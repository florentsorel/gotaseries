package gotaseries

type similarsResponse struct {
	Similars []SimilarShow `json:"similars"`
	Errors   Errors        `json:"errors"`
}

type SimilarShow struct {
	ID        int     `json:"id"`
	Title     string  `json:"show_title"`
	ShowID    int     `json:"show_id"`
	TheTvdbID int     `json:"thetvdb_id"`
	Notes     *string `json:"notes"`
	Show      *Show   `json:"show"`
}
